#/bin/sh

curl -s https://api.github.com/repos/deploy-f/deployf-cli/releases/latest \
| grep 'browser_download_url.*deployf"' \
| cut -d : -f 2,3 \
| tr -d \" \
| xargs -n 1 curl -O -sSL

chmod +x deployf
sudo mv deployf /usr/local/bin