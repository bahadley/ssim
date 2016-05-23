# ssim 


### How can I run it? ###
ssim can be run as a Linux command (for example):     

```
$ SSIM_DST_ADDR=10.0.0.5 ssim    
```

with environment variables configured, or just   

```
$ ssim   
```

to use default values.   


### Configurable environment variables ###

*SSIM_ADDR*=127.0.0.1 (Source IPv4 address for the sensor tuples)      
*SSIM_DST_ADDR*=127.0.0.1 (Comma delimited destination IPv4 addresses for the sensor tuples)    
*SSIM_DST_PORT*=22221 (Destination port for the sensor tuples)    

*SSIM_TRANSMIT*=yes (Enable tracing [yes, no])
*SSIM_TRACE*=no (Enable tracing [yes, no])
