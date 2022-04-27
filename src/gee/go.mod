module gee

go 1.18

require my_utils v1.0.0

replace my_utils => ../my_utils

require golang.org/x/net v0.0.0-20220421235706-1d1ef9303861

replace golang.org/x/net => ../golang.org/x/net v0.0.0-20220421235706-1d1ef9303861