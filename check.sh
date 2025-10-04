#!/bin/bash
PROXY_USERNAME=********
PROXY_PASSWORD=********
SOCKS_PROXY_PORT=1080
HTTPS_PROXY_PORT=8443
arr=(
	"**************"
)

for item in "${arr[@]}"; do

	connect="socks5h://$PROXY_USERNAME:$PROXY_PASSWORD@$item:$SOCKS_PROXY_PORT"
	if curl --max-time 3 -x "$connect" ident.me >/dev/null 2>&1; then
		echo "[OK] $connect"
	else
		echo "[FAILED] $connect"
	fi

	connect="https://$PROXY_USERNAME:$PROXY_PASSWORD@$item:$HTTPS_PROXY_PORT"
	if curl --max-time 3 -x "$connect" --proxy-insecure -k ident.me >/dev/null 2>&1; then
		echo "[OK] $connect"
	else
		echo "[FAILED] $connect"
	fi

done
