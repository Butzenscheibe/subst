# Subst

![subst](./img/subst.png "subst")

__Currently under development and testing. I don't take any responsability for unexpected behavior. You know what you are doing :3__

A simple extension over kustomize, which allows further variable substitution and introduces simplified yet strong secrets management (for multi tenancy use-cases). Extends to functionality of kustomize for ArgoCD users.

# Functionality

The idea for subst is to act as complementary for kustomize. You can reference additional variables for your environment or from different kustomize paths, which are then accesible across your entire kustomize build. The kustomize you are referencing to is resolved (it's paths). In each of these paths you can create new substition files, which contain variables or secrets, which then can be used by your kustomization. The final output is your built kustomization with the substitutions made.

By default the all files are considered using this regex `(subst\\.yaml|vars\\.yaml|.*(ejson|vars))`. You can change the regex using:

```
subst render . --file-regex "custom-values\\.yaml"
```

## Getting Started




## Available Substitutions

You can display which substitutions are available for a kustomize build by running:

```
subst substitutions .
```



See available options with:

```
subst substitutions -h
```

### Paths



### Environment

for environment variables which come from an argo application (`^ARGOCD_ENV_`) we remove the `ARGOCD_ENV_` and they are then available in your substitutions without the `ARGOCD_ENV_` prefix. This way they have the same name you have given them on the application ([Read More](https://argo-cd.readthedocs.io/en/stable/operator-manual/config-management-plugins/#using-environment-variables-in-your-plugin)). All the substions are available as flat key, so where needed you can use environment substitution.

## Spruce

[Spruce](https://github.com/geofffranks/spruce) is used to access the substition variables, it has more flexability than [envsubst](#environment-substitution). You can grab values from the available substitutions using [Spruce Operators](https://github.com/geofffranks/spruce/blob/main/doc/operators.md). Spurce is greate, because it's operators are valid YAML which allows to build the kustomize without any further hacking.


## Environment Substitution

Adds support for [bash string replacement functions](https://github.com/drone/envsubst)

By default environment substitution is not performed. You need to enable environment substituion with the following flag:

```
subst render . --envsubst
```

You can disable environment substitution for a single file by giving it an annotation (assuming `--envsubst` is set):

```
apiVersion: operatingsystemmanager.k8c.io/v1alpha1
kind: OperatingSystemProfile
metadata:
  name: osp-flatcar
  namespace: kube-system
  annotations:
    kustomize.toolkit.fluxcd.io/substitute: "disabled"
....   
```

If you want to skip a single variable, you can use `$${variable}`, this will print `${variable}`

## Secrets

You can both encrypt files which are part of the kustomize build or which are used for substitution. Currently for secret decryption we support both [ejson](https://github.com/Shopify/ejson) and [sops](https://github.com/mozilla/sops). You can use any combination of these decryption providers together. The principal for all decryption provider is, that they should load the private keys while a substiution build is made instead of having a permanent keystore. This allows for secret tenancy (eg. one secret per argo application). 

Decryption can be disabled, in that case the files are just loaded, without their encryption properties (might be useful if you dont have access to the private keys to decrypt the secrets):

```
subst render . --skip-decrypt
```

Decryption can be enforced, if a secret can not be decrypted it's treated as an error (recommended):

```
subst render . --must-decrypt
```

### EJSON

[EJSON](https://github.com/Shopify/ejson) allows simple secrets management. I like it, because you can rencrypt secrets without having the private key, which is sometimes useful. 

You can encrypt entire files using EJSON. The file must be in json format (which is fun for kustomize). The entire file will be encrypted, which may not bes useful in all cases.




### SOPS

[SOPS](https://github.com/mozilla/sops) is commonly known and also used by [FluxCD](https://fluxcd.io/flux/guides/mozilla-sops/). 



### Kubernetes

For all decryptors you can create a kubernetes secret, which contains the private information for secret decryption.






# Running it


## Local installation

**Brew**

```bash
brew tap buttahtoast/subst
```

**Docker**

```bash
docker run -it ghcr.io/buttahtoast/subst -h
```

**Github Releases**

https://github.com/buttahtoast/subst/releases


## ArgoCD Plugin



## CI/CD

TBD


