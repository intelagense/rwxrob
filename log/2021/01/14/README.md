## Thursday, January 14, 2021, 3:01:33AM EST <1610611293>

Just realized that changing my email in `.gitconfig` has caused by
commits to no longer be signed and therefore not show as verified. I
just had to add another identity to the GPG key:

```
$ gpg --edit-key <key-id>
gpg> adduid
Real Name: <name>
Email address: <email>
Comment: <comment or Return to none>
Change (N)ame, (C)omment, (E)mail or (O)kay/(Q)uit? O
Enter passphrase: <password>
gpg> uid <uid>
gpg> trust
Your decision? 5
Do you really want to set this key to ultimate trust? (y/N) y
gpg> save
```
