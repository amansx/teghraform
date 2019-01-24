Feature: Create S3 Buckets if not exist
	This Step Creates S3 Buckets if they don't exist

	@CreateAllBuckets
	Scenario Outline: Seed buckets
		Given DEFINE
		| Name       | Update | Type          |
		| mybucket2  | true   | AWS.S3.Bucket |

		When aws s3 bucket doesn't exist
		| <example> | 123 |

		And mybucket2.Update == 'false'
		
		Then create aws s3 bucket
	
	Examples:
		| Name      | Update | Type          |
		| mybucket  | true   | AWS.S3.Bucket |
		| mybucket1 | false  | AWS.S3.Bucket |

	@Rollback::CreateAllBuckets
	Scenario Outline: Seed buckets
		Given DEFINE
		| Name      | Update | Type          |
		| mybucket  | true   | AWS.S3.Bucket |
		| mybucket1 | false  | AWS.S3.Bucket |

		When aws s3 bucket doesn't exist
		| <example> | mybucket1.Update |

		And mybucket.Update == 'true'

		Then create aws s3 bucket

	Examples:
		| Name      | Update | Type          |
		| mybucket  | true   | AWS.S3.Bucket |
		| mybucket1 | false  | AWS.S3.Bucket |