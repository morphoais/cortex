/*
Copyright 2021 Cortex Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package batchapi

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"gitlab.com/g-ogawa/cortex/pkg/lib/aws"
	"gitlab.com/g-ogawa/cortex/pkg/lib/errors"
	"gitlab.com/g-ogawa/cortex/pkg/operator/config"
	"gitlab.com/g-ogawa/cortex/pkg/operator/schema"
	"github.com/gobwas/glob"
)

// Takes in a function(shouldSkip, bucketName, s3.Object)
func s3IteratorFromLister(s3Lister schema.S3Lister, fn func(string, *s3.Object) (bool, error)) error {
	includeGlobPatterns := make([]glob.Glob, 0, len(s3Lister.Includes))

	for _, includePattern := range s3Lister.Includes {
		globExpression, err := glob.Compile(includePattern, '/')
		if err != nil {
			return errors.Wrap(err, "failed to interpret glob pattern", includePattern)
		}
		includeGlobPatterns = append(includeGlobPatterns, globExpression)
	}

	excludeGlobPatterns := make([]glob.Glob, 0, len(s3Lister.Excludes))
	for _, excludePattern := range s3Lister.Excludes {
		globExpression, err := glob.Compile(excludePattern, '/')
		if err != nil {
			return errors.Wrap(err, "failed to interpret glob pattern", excludePattern)
		}
		excludeGlobPatterns = append(excludeGlobPatterns, globExpression)
	}

	for _, s3Path := range s3Lister.S3Paths {
		bucket, key, err := aws.SplitS3Path(s3Path)
		if err != nil {
			return err
		}

		awsClientForBucket, err := aws.NewFromClientS3Path(s3Path, config.AWS)
		if err != nil {
			return err
		}

		err = awsClientForBucket.S3Iterator(bucket, key, false, nil, func(s3Obj *s3.Object) (bool, error) {
			s3FilePath := aws.S3Path(bucket, *s3Obj.Key)

			shouldSkip := false
			if len(includeGlobPatterns) > 0 {
				shouldSkip = true
				for _, includeGlobPattern := range includeGlobPatterns {
					if includeGlobPattern.Match(s3FilePath) {
						shouldSkip = false
						break
					}
				}
			}

			for _, excludeGlobPattern := range excludeGlobPatterns {
				if excludeGlobPattern.Match(s3FilePath) {
					shouldSkip = true
					break
				}
			}

			if !shouldSkip {
				return fn(bucket, s3Obj)
			}

			return true, nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
