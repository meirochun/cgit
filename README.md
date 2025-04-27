## Git credential manager CLI

This CLI changes the `user.name` and `user.email` from `git config --list`.

You might want to use this if you have a corporative account and a personal one. At least this is why a made this.

## Usage

#### Checks CLI's version:
```
cgit version
```

#### Setting username or email:
```
cgit --username "Your username here" --email "your@email.com"
or
cgit -u "Your username here" -e "your@email.com"
```
<small>If you want to change only the email or username, you also can do it. Just pass the flag you need.</small>