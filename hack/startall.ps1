Start-Job -Name ctl -WorkingDirectory C:\users\Andres\projects\go\jobico\ctl\cmd\ {go run -race main.go}
Start-Sleep -Seconds 4
Start-Job -Name queue -WorkingDirectory C:\users\Andres\projects\go\jobico\srv\cmd\queue\ {go run -race main.go}
Start-Job -Name repo -WorkingDirectory C:\users\Andres\projects\go\jobico\repo\cmd\ {go run -race main.go}
Start-Job -Name recorder -WorkingDirectory C:\users\Andres\projects\go\jobico\recorder\cmd {go run -race main.go}
Start-Job -Name listener -WorkingDirectory C:\users\Andres\projects\go\jobico\srv\cmd\listener\ {go run -race main.go}
Start-Job -Name exec -WorkingDirectory C:\users\Andres\projects\go\jobico\srv\cmd\executor\ {go run -race main.go}
