# ssim (stream simulator)

### What is ssim? ###

It is a simple data stream simulator designed to exercise esp (edge stream processor).

### How can I run it? ###
ssim can be run as a Linux command (for example):     

``` $ SSIM_DST_ADDR=127.0.0.1,127.0.0.2 ssim ```

with environment variables configured, or just   

``` $ ssim ```

to use default values.   


### Configurable environment variables ###

*SSIM_ADDR*=127.0.0.1 (Source IPv4 address for the sensor tuples)      
*SSIM_DST_ADDR*=127.0.0.1 (Comma delimited destination IPv4 addresses for the sensor tuples)    
*SSIM_DST_PORT*=22221 (Destination port for the sensor tuples)    

*SSIM_NUM_TUPLES*=100 (Number of sensor tuples to transmit)    
*SSIM_DELAY_INTERVAL*=100 ms (Number of milliseconds to pause between tuple transmissions)   
*SSIM_AGGREGATE_SIZE*=2 (Number of tuples that will trigger an egress result tuple in esp) 

*SSIM_TRANSMIT*=yes (Enable tracing [yes, no])   
*SSIM_TRACE*=no (Enable tracing [yes, no])   

*SSIM_FLUSH_INTERVAL*=0 (Send a flush tuple per this number of standard tuples)   
*SSIM_FLUSH_DELAY*=10 ms (Number of milliseconds to pause after a flush tuple is sent)   
