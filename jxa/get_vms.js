#!/usr/bin/env osascript -l JavaScript

// Console.log outputs to stderr by default which is SO DUMB
// Workaround to override from here: https://apple.stackexchange.com/a/275811
console.log = function() {
    ObjC.import('Foundation');
    for (argument of arguments) {
        $.NSFileHandle.fileHandleWithStandardOutput.writeData($.NSString.alloc.initWithString(String(argument) + "\n").dataUsingEncoding($.NSNEXTSTEPStringEncoding));
    }
}

const app = Application('UTM');
let vms = app.virtualMachines().map((vm) => {
    return JSON.parse(Automation.getDisplayString(vm.properties()));
});
console.log(JSON.stringify(vms));
