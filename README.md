1. Open the command prompt in the below directory
   transaction_test/tests/
2. To test "Add Transaction" API, type in command console like this:
   go test get_test.go -v
3. To test "List Transactions with pagination and filter" API, type in command console like this:
   go test post_test.go -v

You can change the URL and the post data in each test files, according to your test purpose.
