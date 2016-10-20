# Contributing

Thanks for being interested in contributing! Here's some things you can do to
help.

* Create issues.
* Add a new notification type.
* Add a new trigger.

## Create a notification type

* Create a general notification package in `ntypes`.
* At minimum, you must implement `ntypes.Notification`.

* Create a subcommand in `cli`. (Use existing packages as examples.)
* Add command to `main.go`.

## Create a trigger

* Create a trigger package in `triggers`.
* At minimum, you must implement `triggers.Trigger`.
* Add it to `triggers/run.go`.

## Submission

You must branch off `dev`. I only put working, tested code on `master`.

* Fork Noti on GitHub.
* `cd variadico/noti`
* `git remote add fork git@github.com:YOUR_USERNAME/noti.git`
* `git checkout -b your-feature`
* Make changes
* `git push fork dev`
* Open pull request
