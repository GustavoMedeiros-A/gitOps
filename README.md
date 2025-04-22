# Steps of Kind (k8)

- go install sigs.k8s.io/kind@v0.27.0
- kind create cluster --name=gitops
- kubectl cluster-info --context kind-gitops

Utilização do Kustomize do Kubernete (garantir a mesma versão no repo e no servidor)
- sudo snap install kustomize
- Mudar o sha da newTag e fazer a mudança da imagem
- Então teoricamente o processo de CD é mudar a tag da imagem e comitar

Com esse comando, ele vai mudar a tag do Kustomize e deployar ela
- kustomize edit set image goserver=gutsmedeiros/gitops:newtag