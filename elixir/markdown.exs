defmodule Markdown do
  @doc """
    Parses a given string with Markdown syntax and returns the associated HTML for that string.

    ## Examples

      iex> Markdown.parse("This is a paragraph")
      "<p>This is a paragraph</p>"

      iex> Markdown.parse("# Header!\\n* __Bold Item__\\n* _Italic Item_")
      "<h1>Header!</h1><ul><li><strong>Bold Item</strong></li><li><em>Italic Item</em></li></ul>"
  """
  @spec parse(String.t()) :: String.t()
  def parse(m) do
    String.split(m, "\n")
    |> Enum.map_join(&process_lines/1)
    |> process_tags()
    |> wrap_ul_md()
  end

  defp process_lines(<<"#######", _rest::binary>> = str), do: "<p>#{str}</p>"
  defp process_lines(<<"#", rest::binary>>), do: rest |> parse_header(1)
  defp process_lines(<<"* ", rest::binary>>), do: "<li>#{rest}</li>"
  defp process_lines(str), do: "<p>#{str}</p>"
  defp parse_header(<<" ", rest::binary>>, lvl), do: "<h#{lvl}>#{rest}</h#{lvl}>"
  defp parse_header(<<"#", rest::binary>>, lvl), do: parse_header(rest, lvl + 1)

  defp process_tags(str) do
    str
    |> String.replace(~r/__([^_]+)__/, "<strong>\\1</strong>")
    |> String.replace(~r/_([^_]+)_/, "<em>\\1</em>")
  end

  defp wrap_ul_md(str) do
    String.replace(str, ~r/<li>(.*)<\/li>/, "<ul><li>\\1</li></ul>")
  end
end
