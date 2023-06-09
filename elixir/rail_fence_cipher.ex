defmodule RailFenceCipher do
  @doc """
  Encode a given plaintext to the corresponding rail fence ciphertext
  """
  @spec encode(String.t(), pos_integer) :: String.t()
  def encode(str, rails) do
    matrix = create_matrix(rails, String.length(str))

    encode(str, {0, 0}, 0, rails - 1, :down, matrix)
    |> Enum.reduce("", fn row, acc ->
      Enum.reject(row, fn e -> e == "*" end)
      |> Enum.join()
      |> then(fn res -> "#{acc}#{res}" end)
    end)
  end

  defp encode(str, _coord, start, start, _direction, _matrix), do: [[str]]
  defp encode("", _i, _start, _finish, _direction, matrix), do: matrix

  defp encode(<<ch::binary-size(1), rest::binary>>, {finish, c}, start, finish, :down, matrix) do
    encode(rest, {finish - 1, c + 1}, start, finish, :up, update_matrix(matrix, ch, finish, c))
  end

  defp encode(<<ch::binary-size(1), rest::binary>>, {start, c}, start, finish, :up, matrix) do
    encode(rest, {start + 1, c + 1}, start, finish, :down, update_matrix(matrix, ch, start, c))
  end

  defp encode(<<ch::binary-size(1), rest::binary>>, {r, c}, start, finish, :down, matrix) do
    encode(rest, {r + 1, c + 1}, start, finish, :down, update_matrix(matrix, ch, r, c))
  end

  defp encode(<<ch::binary-size(1), rest::binary>>, {r, c}, start, finish, :up, matrix) do
    encode(rest, {r - 1, c + 1}, start, finish, :up, update_matrix(matrix, ch, r, c))
  end

  defp update_matrix(matrix, ch, r, c) do
    put_in(matrix, [Access.at(r), Access.at(c)], ch)
  end

  defp create_matrix(rows, columns) do
    for _r <- 1..rows, do: for(_c <- 1..columns, do: "*")
  end

  defp collect(_matrix, _coord, _, _, len, acc) when length(acc) == len,
    do: Enum.reverse(acc) |> Enum.join()

  defp collect(matrix, {bound, c}, bound, :down, len, acc) do
    ch = get_in(matrix, [Access.at(bound), Access.at(c)])
    collect(matrix, {bound - 1, c + 1}, bound, :up, len, [ch | acc])
  end

  defp collect(matrix, {0, c}, bound, :up, len, acc) do
    ch = get_in(matrix, [Access.at(0), Access.at(c)])
    collect(matrix, {1, c + 1}, bound, :down, len, [ch | acc])
  end

  defp collect(matrix, {r, c}, bound, :down, len, acc) do
    ch = get_in(matrix, [Access.at(r), Access.at(c)])
    collect(matrix, {r + 1, c + 1}, bound, :down, len, [ch | acc])
  end

  defp collect(matrix, {r, c}, bound, :up, len, acc) do
    ch = get_in(matrix, [Access.at(r), Access.at(c)])
    collect(matrix, {r - 1, c + 1}, bound, :up, len, [ch | acc])
  end

  @doc """
  Decode a given rail fence ciphertext to the corresponding plaintext
  """
  @spec decode(String.t(), pos_integer) :: String.t()
  def decode(str, rails) do
    len = String.length(str)
    matrix = create_matrix(rails, len)

    decode(str, {0, 0}, 0, rails - 1, :down, matrix, matrix |> hd() |> length(), 0)
    |> collect({0, 0}, rails - 1, :down, len, [])
  end

  defp decode(str, _i, start, start, _direction, _matrix, _, _printing_row), do: [[str]]

  defp decode("", _i, _start, _finish, _direction, matrix, _, _printing_row),
    do: matrix

  defp decode(str, {_r, c}, start, finish, _, matrix, columns, printing_row)
       when c > columns - 1 do
    decode(str, {start + 1, start + 1}, start, finish, :down, matrix, columns, printing_row + 1)
  end

  defp decode(
         <<ch::binary-size(1), rest::binary>> = str,
         {finish, c},
         start,
         finish,
         :down,
         matrix,
         columns,
         printing_row
       ) do
    {updated_matrix, updated_str} =
      if finish == printing_row do
        {update_matrix(matrix, ch, printing_row, c), rest}
      else
        {matrix, str}
      end

    decode(
      updated_str,
      {finish - 1, c + 1},
      start,
      finish,
      :up,
      updated_matrix,
      columns,
      printing_row
    )
  end

  defp decode(
         <<ch::binary-size(1), rest::binary>> = str,
         {start, c},
         start,
         finish,
         :up,
         matrix,
         columns,
         printing_row
       ) do
    {updated_matrix, updated_str} =
      if start == printing_row do
        {update_matrix(matrix, ch, printing_row, c), rest}
      else
        {matrix, str}
      end

    decode(
      updated_str,
      {start + 1, c + 1},
      start,
      finish,
      :down,
      updated_matrix,
      columns,
      printing_row
    )
  end

  defp decode(
         <<ch::binary-size(1), rest::binary>> = str,
         {r, c},
         start,
         finish,
         :down,
         matrix,
         columns,
         printing_row
       ) do
    {updated_matrix, updated_str} =
      if r == printing_row do
        {update_matrix(matrix, ch, printing_row, c), rest}
      else
        {matrix, str}
      end

    decode(
      updated_str,
      {r + 1, c + 1},
      start,
      finish,
      :down,
      updated_matrix,
      columns,
      printing_row
    )
  end

  defp decode(
         <<ch::binary-size(1), rest::binary>> = str,
         {r, c},
         start,
         finish,
         :up,
         matrix,
         columns,
         printing_row
       ) do
    {updated_matrix, updated_str} =
      if r == printing_row do
        {update_matrix(matrix, ch, printing_row, c), rest}
      else
        {matrix, str}
      end

    decode(updated_str, {r - 1, c + 1}, start, finish, :up, updated_matrix, columns, printing_row)
  end
end

# msg = "WEAREDISCOVEREDFLEEATONCE"
# msg = "The quick brown fox jumps over the lazy dog."
# msg = "One rail, only one rail"
# RailFenceCipher.encode("XOXOXOXOXOXOXOXOXO", 2)
# RailFenceCipher.encode(msg, 1)
# cipher = "TEITELHDVLSNHDTISEIIEA"
# cipher = "WECRLTEERDSOEEFEAOCAIVDEN"
cipher = "TEITELHDVLSNHDTISEIIEA"
# cipher = "EIEXMSMESAORIWSCE"

RailFenceCipher.decode(cipher, 3)
|> IO.inspect()

# "EXERCISMISAWESOME"
# "THEDEVILISINTHEDETAILS"
# "WECRLTEERDSOEEFEAOCAIVDEN"
