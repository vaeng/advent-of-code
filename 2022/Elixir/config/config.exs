import Config

config :advent_of_code, AdventOfCode.Input,
  allow_network?: true,
  session_cookie: "53616c7465645f5fc201c4cfa270a32999c2034f55723d97cc5f601ee27efaeb71fdc015502aa08b14dfc0277160b272a6b6f30fff8d53bf688c14b593f9b930"

# If you don't like environment variables, put your cookie in
# a `config/secrets.exs` file like this:
#
# use Mix.Config
# config :advent_of_code, AdventOfCode.Input,
#   session_cookie: "..."

try do
  import_config "secrets.exs"
rescue
  _ -> :ok
end
