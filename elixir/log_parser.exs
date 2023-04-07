defmodule LogParser do
  def valid_line?(line) do
    String.match?(line, ~r/^\S+((ERROR)|(INFO)|(WARNING)|(DEBUG))/)
  end

  def split_line(line) do
    String.split(line, ~r/\<([\~\*\=\-])*\>/)
  end

  def remove_artifacts(line) do
    String.replace(line, ~r/(end-of-line)\S*[0-9]/i, "")
  end

  def tag_with_user_name(line) do
    String.match?(line, ~r/(User)\s/i)
    |> add_tag(line)
  end

  defp add_tag(match?, line) when match? do
    [name] = Regex.run(~r/(?<=User)\s+(\S+)/i, line, capture: :first)
    trimed_name = String.trim(name)

    "[USER] #{trimed_name} #{line}"
  end

  defp add_tag(_, line), do: line
end
