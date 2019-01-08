@seedS3
Feature: Create S3 Buckets if not exist
	This Step Creates S3 Buckets if they don't exist

	Scenario: seed buckets
		Given AWS.S3.Bucket
		| Name      | Update |
		| mybucket  | true   |
		| mybucket1 | false  |

		When aws s3 bucket doesn't exist
		| mybucket |
		
		And mybucket.Update == 'true'
		
		Then create aws s3 bucket
		| mybucket |	
