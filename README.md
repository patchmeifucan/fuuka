# Fuuka - The Threat Hunting Filewalker
![Fuuka Banner](https://github.com/patchmeifucan/fuuka/blob/main/fuuka_banner.png)

Fuuka is a tool written in Go that matches readable file data to a YARA ruleset in order to attempt to detect Malware and other IOCs. As of currently, Fuuka only performs scans against files to match against a compiled YARA ruleset. It is not intended to be a silver bullet solution, but rather a way to quickly enumerate a file system to find potential trails of compromise on a target system.

## Usage
For sake of performance, Fuuka will only load compiled YARA rulesets. You can compile a YARA ruleset from the command line with `yarac in.yar out` and it's ready for usage.

Syntax:
`fuuka <-p/--path> [scan path] <-j/--jobs> [jobs] --yara [YARA ruleset]`

## Requirements
<ul>
<li>go</li>
<li>yara</li>
</ul>

## Authors
[patchmeifucan](https://github.com/patchmeifucan)<br>
[haven7](https://github.com/HeavenSmiles)<br>

## License
Fuuka is licensed under the AGPL 3.0, you can read the license [here](https://github.com/patchmeifucan/fuuka/blob/main/LICENSE)
