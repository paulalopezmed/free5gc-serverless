<!-- # Serverless 5G Code



## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

- [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
- [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin https://git.tu-berlin.de/master-theses/serverless-5g-code.git
git branch -M main
git push -uf origin main
```

## Integrate with your tools

- [ ] [Set up project integrations](https://git.tu-berlin.de/master-theses/serverless-5g-code/-/settings/integrations)

## Collaborate with your team

- [ ] [Invite team members and collaborators](https://docs.gitlab.com/ee/user/project/members/)
- [ ] [Create a new merge request](https://docs.gitlab.com/ee/user/project/merge_requests/creating_merge_requests.html)
- [ ] [Automatically close issues from merge requests](https://docs.gitlab.com/ee/user/project/issues/managing_issues.html#closing-issues-automatically)
- [ ] [Enable merge request approvals](https://docs.gitlab.com/ee/user/project/merge_requests/approvals/)
- [ ] [Automatically merge when pipeline succeeds](https://docs.gitlab.com/ee/user/project/merge_requests/merge_when_pipeline_succeeds.html)

## Test and Deploy

Use the built-in continuous integration in GitLab.

- [ ] [Get started with GitLab CI/CD](https://docs.gitlab.com/ee/ci/quick_start/index.html)
- [ ] [Analyze your code for known vulnerabilities with Static Application Security Testing(SAST)](https://docs.gitlab.com/ee/user/application_security/sast/)
- [ ] [Deploy to Kubernetes, Amazon EC2, or Amazon ECS using Auto Deploy](https://docs.gitlab.com/ee/topics/autodevops/requirements.html)
- [ ] [Use pull-based deployments for improved Kubernetes management](https://docs.gitlab.com/ee/user/clusters/agent/)
- [ ] [Set up protected environments](https://docs.gitlab.com/ee/ci/environments/protected_environments.html)

***

# Editing this README

When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!). Thank you to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README
Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a self-explaining name for your project.

## Description
Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your project, this is a good place to list differentiating factors.

## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals
Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation
Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support
Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For open source projects, say how it is licensed.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers. -->


## For installing and connect to openfaas dashboard

1. Install minikube and start it
2. Install with arkade openfaas and verify it works with: To verify that openfaas has started, run:
```bash
kubectl -n openfaas get deployments -l "release=openfaas, app=openfaas"
```
3. To know my admin password decoded run: 
 ```console
 echo $(kubectl -n openfaas get secret basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode)

fTZsLHZgOahZ
```
4.  Forward the gateway to your machine
```bash
kubectl rollout status -n openfaas deploy/gateway
kubectl port-forward -n openfaas svc/gateway 8080:8080 &
```
In case it gives and error it is already in use, try with killing the process by geeting first the ID with:

```bash
ps aux | grep kubectl
```
And then kill the process with: 
```bash
kill ID
```

5. For getting in the openfaas dashboard

    5.1. Lynx:
    ```bash 
    lynx http://127.0.0.1:8080/ui/
    ```
    5.2. Curl:
    ```bash
    curl -u admin:fTZsLHZgOahZ http://127.0.0.1:8080/ui/
    ```


6. You don't need the web dashboard to use OpenFaaS, you can get the CLI with this command: curl -sL cli.openfaas.com | sh

7. For making alias k=kubectl in ~./bashrc





## Access to the openfaas cluster throught and Ubuntu 20.04.6 WSL 

Following this tutorial (https://www.zepworks.com/posts/access-minikube-remotely-kvm/#3c-open-the-port-)

1. Copy the private key of the server machine to the WSL in ~/.ssh/id_rsa

2. Copy all the certificates from minikube in WSL
```bash
	scp user@IP:ca.crt .
	scp user@IP:client.crt .
	scp user@IP:client.key .
```

3. Install nginx in the server machine

```bash
	sudo apt update
	sudo apt install nginx
```

4. In /etc/nginx/nginx.conf add this piece of configuration
This line "listen <IP>:8080;" refers to the IP of the server machine
```bash 	
stream {
  server {
      # This is the Public IP and Port
      listen <IP>:8080;
      #TCP traffic will be forwarded to the specified server
      # This is the Minikube IP and Port
      proxy_pass 192.168.49.2:8443;
  }
}

```

5. And restart the nginx:
```bash
	sudo systemctl restart nginx
```
6. In /.kube/ create a file called config with this text:
In the line "server: https://<IP>:8080" we are refering to the previos config we created in the server machine in the nginx configuration
```
apiVersion: v1
clusters:
- cluster:
    certificate-authority: ca.ctr
    # server: https://192.168.39.131:8443  <-- this was the previous value there
    server: https://<IP>:8080"
  name: minikube
contexts:
- context:
    cluster: minikube
    user: <username>
  name: minikube
current-context: minikube
kind: Config
preferences: {}
users:
- name: <username>
  user:
    client-certificate: client.ctr
    client-key: client.key
```
For checking  conection with the cluster:

```bash
kebectl get pods
```

7. Continue with the previous set:

    7.1. First get the admin password: 
    ```bash
    echo $(kubectl -n openfaas get secret basic-auth -o jsonpath="{.data.basic-auth-password}" | base64 --decode)
    ```
    7.2. Forward the 8080 port:
    ```bash
     kubectl port-forward -n openfaas svc/gateway 8080:8080 &
     ```
    7.3. Check the openfaas cluster:
    ```bash
     kubectl -n openfaas get deployments -l "release=openfaas, app=openfaas"
     ```
    7.4. Access the openfaas dashboard:
    ```bash
     curl -u admin:<password> http://127.0.0.1:8080/ui/
    ```


----------------------------------
Synchronize the deployment of the cluster acording the new helm configuration (like changing admin password):
```bash 
helm upgrade openfaas --install chart/openfaas --namespace openfaas -f ./chart/openfaas/values.yaml
``` 

## For deploying a hello world function in openfaas


1. ERROR: docker image list
Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock: Get "http://%2Fvar%2Frun%2Fdocker.sock/v1.24/images/json": dial unix /var/run/docker.sock: connect: permission denied

	For fixing this error we have to add the user docker to the list  of users:
    ```bash 
	sudo usermod -aG docker $USER
	newgrp docker		
    ```
    Make sure that docker is in the list of sudo users:
    ```bash
    groups
    ```

2. Download the go template:
    ```bash
     faas-cli new --lang go gohash
    ```

3. Edit the gohash.yml file which was generated by the CLI. Put your Docker Hub account into the image: section. i.e. image: paulalopezmed/gohash:latest.

4. 
```bash
 faas-cli build -f gohash.yml
```
5. 
```bash
 faas-cli push -f gohash.yml
```

6. 
```bash
 faas-cli deploy -f gohash.yml
```

7. For testing: 
```bash
echo -n "test" | faas-cli invoke gohash
```

## 5G process selection: NAS Authentication process 
For the purpose of simplifying we have choosen the process for the following criteria:
1) First AMF operation when connecting a UE. With that it will be easier to compare the functionality by measuring the performace at connecting the UE.(https://www.sharetechnote.com/html/OpenRAN/OR_free5GC_Run.bak)
2) Insolated process that only involves the AMF and not any 5gcore function
(3) Easily insolated code out of the repository. Clear Encode and Decode functions, and any functions that refers them.)  

NAS Authentication process steps (PAGE 54 https://www.etsi.org/deliver/etsi_ts/133500_133599/133501/15.02.00_60/ts_133501v150200p.pdf)
(PAGE 41 https://www.etsi.org/deliver/etsi_ts/124500_124599/124501/16.05.01_60/ts_124501v160501p.pdf):

1) Authentication Request: when a UE initiates the connection to the 5G network, it sends an Authentication Request message to the AMF. This message contains the UE's identity.
2) Identity Verification: The AMF recieves the Authentication Request message and verifies the UE's identity.
3) NAS Security Mode Command (HASH): After successful authentication of the UE, the AMF sends a security mode command that includes information about the security algorithms and keys to be used for securing the communications between the UE and the network
4) NAS Security Mode Complete: indicates that the UE has completed the establishment of the requested security modes
5) NAS Registration Complete: notify the AMF that it has completed the registration process. UE ready for using the 5g network resources.


