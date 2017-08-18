Get into the container
```bash
docker run -it --rm -v $(pwd):/code -w /code elixir bash
```

Do what you must.

### NOTE
To properly run test, add this to your test file:
```elixir
Code.load_file("rock_paper_scissors.exs", __DIR__)

ExUnit.start
ExUnit.configure trace: true, exclude: :pending
```
