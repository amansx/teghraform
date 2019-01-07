@seedS3
Feature: Create S3 Buckets if not exist
	This Step Creates S3 Buckets if they don't exist

	Scenario: seed buckets
		Given aws.s3.bucket
		| Name      | Update |
		| mybucket  | true   |
		| mybucket1 | false  |

		When aws s3 bucket doesn't exist "bucket1"
		
		And mybucket.Update IS true
		
		Then create aws s3 bucket "bucket1"