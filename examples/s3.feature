@seedS3
Feature: Create S3 Buckets if not exist
	This Step Creates S3 Buckets if they don't exist

	Scenario: seed buckets
		Given AWS.S3.Bucket
		| Name      |
		| mybucket  |
		| mybucket1 |

		When aws s3 bucket doesn't exist
		| mybucket |
		
		And mybucket.Update IS true
		
		Then create aws s3 bucket
		| mybucket |