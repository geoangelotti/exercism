class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    @line.split(": ")[1].strip
  end

  def log_level
    start_index = @line.index "["
    end_index = @line.index "]"
    @line.slice(start_index + 1, end_index-1).downcase
  end

  def reformat
    "#{self.message} (#{self.log_level})"
  end
end
