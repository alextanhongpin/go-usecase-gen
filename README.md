
create a new blog account

| Title | Description |
| -     | -           |
| Name | create a new blog account |
| Related requirements | Requirement A.1<br/>Requirement A.2|
| Goal in Context | A new existing author requests a new blog account from the Administrator|
| Preconditions | The author has appropriate proof of identity.|
| Successful end conditions | a new blog account is created for the author.|
| Failed end conditions | The application for a new blog account is rejected.|
| Primary Actors | Administrator|
| Secondary Actors | Author Credential Database|
| Trigger | The administrator asks the CMS to create a new blog account|
| Main flow | 1. The Administrator asks the system to create a new blog<br/>2. The Administrator selects an account type.<br/>3. The Administrator enters the author's details.<br/>4. The author's details are verified using the Author Credential Database.<br/>5. The new blog account is created.<br/>6. A summary of the new blog account's details are emailed to the author.|
| Extensions | 4.1. The Author Credentials Database does not verify the Author's details.<br/>4.2. The author's blog account application is rejected.|
