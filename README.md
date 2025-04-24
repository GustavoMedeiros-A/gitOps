# Steps of Kind (k8)

- go install sigs.k8s.io/kind@v0.27.0
- kind create cluster --name=gitops
- kubectl cluster-info --context kind-gitops

# Utilização do Kustomize do Kubernete (garantir a mesma versão no repo e no servidor)
- sudo snap install kustomize
- Mudar o sha da newTag e fazer a mudança da imagem
- Então teoricamente o processo de CD é mudar a tag da imagem e comitar

# Com esse comando, ele vai mudar a tag do Kustomize e deployar ela
- kustomize edit set image goserver=gutsmedeiros/gitops:newtag

# Kubectl
- kubectl get deploy
- kubectl get svc

# Instalar o Argo CD
- kubectl create namespace argocd
- kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml


- kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 -d && echo
    - gSPezp6spC-MXEFv
- kubectl port-forward svc/argocd-server -n argocd 8082:443 (Com UI)

- Acessar localhost:8082 (nesse caso)
    - Username: Admin
    - password: gSPezp6spC-MXEFv (obs: gerada nesse caso)