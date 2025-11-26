#!/bin/bash

# Wehrmacht Encyclopedia - Kubernetes Deployment
# Target: 34.29.127.41:7000

set -e

echo "🚀 Wehrmacht Encyclopedia - Kubernetes Deployment"
echo "=================================================="
echo ""

# Check kubectl
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl bulunamadı!"
    exit 1
fi

echo "✅ kubectl bulundu"

# Deploy
echo ""
echo "Deploying..."
echo ""

echo "1️⃣  Deployment oluşturuluyor..."
kubectl apply -f k8s/deployment.yaml

echo "2️⃣  Service oluşturuluyor (NodePort 30700)..."
kubectl apply -f k8s/service.yaml

echo "3️⃣  Ingress oluşturuluyor (opsiyonel)..."
kubectl apply -f k8s/ingress.yaml 2>/dev/null || echo "   ⚠️  Ingress Controller yok, atlanıyor..."

# Wait
echo ""
echo "⏳ Deployment hazırlanıyor..."
kubectl wait --for=condition=available --timeout=300s deployment/wehrmacht-app

# Status
echo ""
echo "📊 Durum:"
kubectl get deployment,service -l app=wehrmacht
kubectl get ingress -l app=wehrmacht 2>/dev/null || true

echo ""
echo "📝 Pods:"
kubectl get pods -l app=wehrmacht

echo ""
echo "✅ Deployment tamamlandı!"
echo ""
echo "🌐 Erişim Seçenekleri:"
echo "====================="
echo ""
echo "1️⃣  NodePort (Port Forwarding ile):"
echo "   Bulut sunucuda çalıştır:"
echo "   sudo iptables -t nat -A PREROUTING -p tcp --dport 7000 -j REDIRECT --to-port 30700"
echo "   Erişim: http://34.29.127.41:7000"
echo ""
echo "2️⃣  Ingress (Ingress Controller varsa):"
echo "   Erişim: http://34.29.127.41.nip.io"
echo "   veya domain ile: http://wehrmacht.example.com"
echo ""
echo "📋 Faydalı Komutlar:"
echo "  Loglar:    kubectl logs -l app=wehrmacht -f"
echo "  Pods:      kubectl get pods -l app=wehrmacht"
echo "  Restart:   kubectl rollout restart deployment/wehrmacht-app"
echo "  Sil:       kubectl delete -f k8s/"
echo ""
