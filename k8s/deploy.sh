#!/bin/bash

# Wehrmacht Encyclopedia - Kubernetes Deployment
# Ingress ile deployment (host-based routing)

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

echo "2️⃣  Service oluşturuluyor (ClusterIP)..."
kubectl apply -f k8s/service.yaml

echo "3️⃣  Ingress oluşturuluyor..."
kubectl apply -f k8s/ingress.yaml

# Wait
echo ""
echo "⏳ Deployment hazırlanıyor..."
kubectl wait --for=condition=available --timeout=300s deployment/wehrmacht-app

# Status
echo ""
echo "📊 Durum:"
kubectl get deployment,service,ingress -l app=wehrmacht

echo ""
echo "📝 Pods:"
kubectl get pods -l app=wehrmacht

echo ""
echo "✅ Deployment tamamlandı!"
echo ""
echo "🌐 Erişim:"
echo "=========="
echo ""
echo "URL:  http://wehrmacht.34.29.127.41.nip.io"
echo "Port: 80 (Ingress)"
echo ""
echo "💡 Not: Başka uygulamanız ile aynı port 80'i paylaşıyor,"
echo "    farklı subdomain ile ayrılıyor (host-based routing)"
echo ""
echo "📋 Faydalı Komutlar:"
echo "  Loglar:    kubectl logs -l app=wehrmacht -f"
echo "  Pods:      kubectl get pods -l app=wehrmacht"
echo "  Ingress:   kubectl get ingress"
echo "  Restart:   kubectl rollout restart deployment/wehrmacht-app"
echo "  Sil:       kubectl delete -f k8s/"
echo ""
