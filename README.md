Test case illustrating a potential issue with the Firestore client

To run it, start the Firestore emulator (`gcloud emulators firestore start`),
then in another shell set the `FIRESTORE_EMULATOR_HOST` environment variable as it suggests.
Then run `go test firestore_test.go`, you should see an assertion failure from line 28,
which I wouldn't expect in this case.
