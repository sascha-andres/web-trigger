# web-trigger #

## Configuration ##

Config file can be passed as the single parameter. If not the current path with a 
value of web-trigger.yml will be used

### Configuration values ###

A sample configuration looks like this:

    ---

    adress: ":8080"
    trigger:
      - route: "test"
        executable: "echo hello"
        checkExecutable: false

web-trigger will generate two routes for this config file:

`/test/trigger` will execute `echo hello` and `/test/log` will show the output of the last call for `/test/trigger`

You may add multiple triggers
