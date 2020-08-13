# log-mover

a tool for moving logback rolling log files to another directory

usage:

```bash
./log-mover -srcDir=/your/log/directory/ -destDir=/your/destination/directory/ -appName=your-application-name -suffix=.log.gz(optional)
``` 

use crontab to move log periodically:

```bash
0 0 * * * log-mover -srcDir=/your/log/directory/ -destDir=/your/destination/directory/ -appName=your-application-name -suffix=.log.gz(optional)
```