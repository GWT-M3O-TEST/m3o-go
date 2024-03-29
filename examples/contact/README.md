# Contact

An [m3o.com](https://m3o.com) API. For example usage see [m3o.com/contact/api](https://m3o.com/contact/api).

Endpoints:

## Update

Update a contact


[https://m3o.com/contact/api#Update](https://m3o.com/contact/api#Update)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/contact"
)

// Update a contact
func UpdateAcontact() {
	contactService := contact.NewContactService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := contactService.Update(&contact.UpdateRequest{
		Links: []contact.Link{
contact.Link: {
Label: "blog", Url: "https://blog.joe.me", },
},
Birthday: "1995-01-01",
Note: "this person is very important",
Id: "42e48a3c-6221-11ec-96d2-acde48001122",
Name: "joe",
Addresses: []contact.Address{
contact.Address: {
Label: "company address", Location: "123 street address", },
},
SocialMedias: []contact.SocialMedia{
contact.SocialMedia: {
Label: "facebook", Username: "joe-facebook", },
},
Phones: []contact.Phone{
contact.Phone: {
Label: "work", Number: "010-87654321", },
},
Emails: []contact.Email{
contact.Email: {
Label: "work", Address: "work@example.com", },
},
	})
	fmt.Println(rsp, err)
	
}
```
## Read

Read contact details


[https://m3o.com/contact/api#Read](https://m3o.com/contact/api#Read)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/contact"
)

// Read contact details
func GetAcontact() {
	contactService := contact.NewContactService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := contactService.Read(&contact.ReadRequest{
		Id: "42e48a3c-6221-11ec-96d2-acde48001122",
	})
	fmt.Println(rsp, err)
	
}
```
## Delete

Delete a contact


[https://m3o.com/contact/api#Delete](https://m3o.com/contact/api#Delete)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/contact"
)

// Delete a contact
func DeleteAcontact() {
	contactService := contact.NewContactService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := contactService.Delete(&contact.DeleteRequest{
		Id: "42e48a3c-6221-11ec-96d2-acde48001122",
	})
	fmt.Println(rsp, err)
	
}
```
## List

List contacts


[https://m3o.com/contact/api#List](https://m3o.com/contact/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/contact"
)

// List contacts
func ListContactsWithDefaultOffsetAndLimitDefaultLimitIs20() {
	contactService := contact.NewContactService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := contactService.List(&contact.ListRequest{
		
	})
	fmt.Println(rsp, err)
	
}
```
## List

List contacts


[https://m3o.com/contact/api#List](https://m3o.com/contact/api#List)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/contact"
)

// List contacts
func ListContactsWithSpecificOffsetAndLimit() {
	contactService := contact.NewContactService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := contactService.List(&contact.ListRequest{
		Offset: 1,
Limit: 1,
	})
	fmt.Println(rsp, err)
	
}
```
## Create

Create a contact


[https://m3o.com/contact/api#Create](https://m3o.com/contact/api#Create)

```go
package example

import(
	"fmt"
	"os"

	"go.m3o.com/contact"
)

// Create a contact
func CreateAcontact() {
	contactService := contact.NewContactService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := contactService.Create(&contact.CreateRequest{
		Note: "this person is very important",
Name: "joe",
Phones: []contact.Phone{
contact.Phone: {
Label: "work", Number: "010-87654321", },
},
Emails: []contact.Email{
contact.Email: {
Address: "work@example.com", Label: "work", },
},
Links: []contact.Link{
contact.Link: {
Label: "blog", Url: "https://blog.joe.me", },
},
Birthday: "1995-01-01",
Addresses: []contact.Address{
contact.Address: {
Label: "company address", Location: "123 street address", },
},
SocialMedias: []contact.SocialMedia{
contact.SocialMedia: {
Label: "facebook", Username: "joe-facebook", },
},
	})
	fmt.Println(rsp, err)
	
}
```
