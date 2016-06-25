# web-trigger #

This is a small rest server written in go that can execute binaries on the machine
it runs on. The output may then be requested on a special url. 

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

## Result values ##

### Not existing routes ###

Return: 404

### Trigger ###

The following JSON is defined as a result:

    {
      status: true,
      msg: ""
    }

`status` will be false when `checkExecutable` is true and the file is not found

When `status` is true, the Process was started

### Log ###

The following JSON is defined as a result:

    {
      status: true,
      log: "some text"
    }

If the log entry does not exist, a 404 will be returned.

### Other ###

If the routes are called in a wrong waz a 500 error is returned

## History ##

### TBD ###