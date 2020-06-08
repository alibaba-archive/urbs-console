# Change Log

All notable changes to this project will be documented in this file starting from version **v1.0.0**.
This project adheres to [Semantic Versioning](http://semver.org/).

-----
## [v0.6.0] - 2020-06-10

**Change:**
- Add API `DELETE /v1/products/{product}/modules/{module}/settings/{setting}:cleanup` that cleanup all rules, users and groups on the setting.
- Add API `DELETE /v1/products/{product}/labels/{label}:cleanup` that cleanup all rules, users and groups on the label.
- Fix recall.
- Update hook event. 


## [v0.5.0] - 2020-06-03

**Change:**

- Support hook.
- Fix search.
- Remove delete label button.

## [v0.4.0] - 2020-05-26

**Change:**

- Add ac API.
- Add test on github.
- Add lock for sync group members.
- Add ac UI.
- Support more query for user and group' settings API.