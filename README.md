# terracounts
Operator for using terraform in multiple accounts at once when deploying code

## Data File required

This app looks for either a file you pass in args or a local data.json file.
This file is JSON formated.
Current structure is a list of strings only containing the account number.

```json
[
  "12345678",

  "98765432"
]
```
If you have feature requests let me know and we can work to add them in.

## Required Terraform Variable use.

You must have account variable existed and used to run assume role like the example below or in a MAP object variable with the account number as its own string item to reference to and set the API keys and secrets dynamically.

here some working examples

```
variable "account" {
  type = string
  default = ""
}
```

Here is the assume role setup with a dynamic map

```
provider "aws" {
  assume_role {
    role_arn = "arn:aws:iam::${var.account}:role/${lookup(element([for i in var.aws_accounts: i if i.account == var.account], 0), "role_name", null)}"  
  }
}
```

## CHANGELOG fun stuf

works with terraform 1.1.1 and aws provider <= 4.13

Written in Go 1.18