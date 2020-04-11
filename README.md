# Go templates
Two packages:
- text/template
- html/template

The `html/template` does the same as the text version except that it understands the context in which data which data will be injected into the template.
To avoid script injection, and resolve common issues such as have to encode special characters for URLs

