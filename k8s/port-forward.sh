#!/bin/bash

# Wehrmacht Encyclopedia - Port 7000 Forwarding Setup
# Bu script bulut sunucuda çalıştırılmalı

echo "🔧 Port 7000 → 30700 Forwarding Kurulumu"
echo "========================================="
echo ""

# Check if running as root
if [ "$EUID" -ne 0 ]; then 
    echo "❌ Bu script root olarak çalıştırılmalı!"
    echo "Şunu dene: sudo ./k8s/port-forward.sh"
    exit 1
fi

echo "✅ Root yetkisi var"
echo ""

# Check if iptables exists
if ! command -v iptables &> /dev/null; then
    echo "❌ iptables bulunamadı!"
    exit 1
fi

echo "✅ iptables bulundu"
echo ""

# Check if rule already exists
if iptables -t nat -L PREROUTING -n | grep -q "dpt:7000"; then
    echo "⚠️  Port 7000 forwarding kuralı zaten var!"
    echo ""
    echo "Mevcut kuralı görmek için:"
    echo "  sudo iptables -t nat -L PREROUTING -n -v | grep 7000"
    echo ""
    echo "Kuralı silmek için:"
    echo "  sudo iptables -t nat -D PREROUTING -p tcp --dport 7000 -j REDIRECT --to-port 30700"
    echo ""
    exit 0
fi

# Add forwarding rule
echo "📝 Port forwarding kuralı ekleniyor..."
iptables -t nat -A PREROUTING -p tcp --dport 7000 -j REDIRECT --to-port 30700

if [ $? -eq 0 ]; then
    echo "✅ Kural eklendi!"
else
    echo "❌ Kural eklenemedi!"
    exit 1
fi

echo ""
echo "📋 Mevcut kural:"
iptables -t nat -L PREROUTING -n -v | grep 7000

echo ""
echo "💾 Kuralı kalıcı yapmak için:"
echo "================================"
echo ""

# Check OS and suggest appropriate command
if [ -f /etc/debian_version ]; then
    echo "Debian/Ubuntu için:"
    echo "  sudo apt-get install iptables-persistent"
    echo "  sudo netfilter-persistent save"
elif [ -f /etc/redhat-release ]; then
    echo "RHEL/CentOS için:"
    echo "  sudo service iptables save"
else
    echo "  sudo iptables-save > /etc/iptables/rules.v4"
fi

echo ""
echo "✅ Port forwarding hazır!"
echo ""
echo "🌐 Erişim: http://34.29.127.41:7000"
echo ""
