It occurred me while talking through the difference between the content
and the rendering that the YAML data inside any component should *never*
include rendering hints of any kind. That stuff should go in
tool-specific configuration files that have render-specific parameters.
For example in the `.kn` directory within any node or th over knowledge
base.
