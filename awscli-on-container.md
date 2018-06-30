I prepared `aws-tools-container` .  
Please clone [here](https://github.com/ygnmhdtt/aws-tools-container).

日本語テスト

---

# Credential

For `aws` command, we need to prepare `~/.aws/config` and `~/.aws/credentials` .  
`aws-tools-container` mounts your local `~/.aws` directory to container home directory.  
So, you don't have to fix them for aws-tools-container.

---

# Usage

Clone repository, and

```
$ make build
$ make up
$ make login

# Now, you are in container

# aws s3 ls
```

---

Default AWS credential is set `default` .  
If you want to change it, please do like this:

```
$ export AWS_DEFAULT_PROFILE=your-profile
```

I prepared profile list command.

```
$ make list
```

---

And, this includes [aws-shell](https://github.com/awslabs/aws-shell) .

```
$ make login
# aws-shell
```

---

If you want to stop container, just type this

```
$ make stop
```

Please use it!!
Now we no longer need pip!!
