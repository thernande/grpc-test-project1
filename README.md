# grpc-todo-app

commands
buf generate proto
bazelisk run //path:command
gazelle
gazelle update-repos -from_file="go.work" -to_macro="deps.bzl%go_dependencies" -prune -build_file_proto_mode=disable_global