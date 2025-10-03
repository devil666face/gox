#!/bin/bash
PROXY_USERNAME=********
PROXY_PASSWORD=********
PROXY_PORT=31080
arr=(
	"**************"
)
for item in "${arr[@]}"; do
	connect="socks5h://$PROXY_USERNAME:$PROXY_PASSWORD@$item:$PROXY_PORT"
	if curl --max-time 1 -x "$connect" ident.me >/dev/null 2>&1; then
		echo "[OK] $connect"
	else
		echo "[FAILED] $item"
	fi
done
