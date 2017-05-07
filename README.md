# intel_backlight_ctl
A simple program to change the backlight level

## Usage 

```
go get github.com/irth/intel_backlight_ctl
sudo chown root $GOPATH/bin/intel_backlight_ctl
sudo chmod u+s $GOPATH/bin/intel_backlight_ctl

intel_backlight_ctl         # prints the current brightness (in percent)
intel_backlight_ctl 15      # adds 15 to the brightness (20% becomes 35%)
intel_backlight_ctl -2      # substracts 2
intel_backlight_ctl set 50  # sets the brightness to 50 percent
```

## Why?
`xbacklight` does not work for me.  
And I can't setuid shellscripts.
