# frozen_string_literal: true
require 'inspec'

CONFIG_INPUT = "1,A,2,B,C\n1,A,2,B,C\n1,A,2,B,C\n1,A,2,B,C\n1,A,2,B,C\n1,A,2,B,C\n1,A,2,B,C\n1,A,2,B,C\n1,A,2,B,C\n1,A,2,B,C"

describe command("echo \"#{CONFIG_INPUT}\" | ./gridgo configure -") do
  its('exit_status') { should eq 0 }
  its('stderr') { should cmp '' }
end

describe file('grid.txt') do
  its('content') { should match 'A,1,A,2,B,C' }
end

describe command("echo \"#{CONFIG_INPUT}\" | ./gridgo configure -f another_grid.txt -") do
  its('exit_status') { should eq 0 }
  its('stderr') { should cmp '' }
end

describe file('another_grid.txt') do
  its('content') { should match 'A,1,A,2,B,C' }
end

describe command('./gridgo "[A3] [D5] [G2]"') do
  its('exit_status') { should eq 0 }
  its('stderr') { should cmp '' }
  its('stdout') { should match '2CA'}
end

describe command('./gridgo -f sample_grid.txt "[A3] [D5] [G2]"') do
  its('exit_status') { should eq 0 }
  its('stderr') { should cmp '' }
  its('stdout') { should match '352'}
end
