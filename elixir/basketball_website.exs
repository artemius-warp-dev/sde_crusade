defmodule BasketballWebsite do
  def extract_from_path(data, path) do
    keys = String.split(path, ".")
    deep = fn
      data, [k | []], _f -> data[k]
      data, [k | tail], f -> f.(data[k], tail, f)
    end

    deep.(data, keys, deep)
  end

  def get_in_path(data, path) do
    keys = String.split(path, ".")
    get_in(data, keys)
  end
end

