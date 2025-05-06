# Тестовое задание YADRO Импульс 2025
## Инструкция по запуску программы
Результат работы программы происходит в файлы (по умолчанию `output.txt` и `result.txt`) и в консоль
#### Для изменения параметров ввода и вывода программы используются следующие флаги:
- `-config` - используется для указания пути к JSON-файлу конфигурации
- `-input` - используется для указания пути к файлу входящих событий
- `-output` - используется для указания пути к создаваемому лог-файлу событий
- `-result` - используется для указания пути к файлу, хранящему результаты соревнований

#### Команда для запуска программы с флагами выглядит следующим образом:
```
 go build main.go
 go run main.go -config="sunny_5_skiers/config.json" -input="sunny_5_skiers/events" -output="output.txt" -result="result.txt" 
 ```
#### Для запуска на предустановленных параметрах:
```
go build main.go
go run main.go
```
### Результаты выводятся в виде строки значений формата:
```
[00:26:22.472] 5 [ {00:13:20.939, 4.375}, {00:13:01.202, 4.481}] {{00:01:40.000, 1.500}, {00:00:50.000, 3.000}} 7/10 
```
```
[Not finished] 1 [ {00:29:02.867, 2.096}] {{00:01:52.476, 0.446}} 4/5
```
#### - где первая ячейка может иметь следующие статусы согласно описанию системы ниже:
- `Not started` - участник дисквалифицирован
- `Not finished` - участник не может продолжать заезд
- `Finished` - отображается как длительность всего заезда
### Для предоставленных исходных данных был получен следующий вывод:
`output.txt` (файл логов)
```
[09:31:49.285] The competitor(3) registered
[09:32:17.531] The competitor(2) registered
[09:37:47.892] The competitor(5) registered
[09:38:28.673] The competitor(1) registered
[09:39:25.079] The competitor(4) registered
[09:55:00.000] The start time for the competitor(1) was set by a draw to 10:00:00.000
[09:56:30.000] The start time for the competitor(2) was set by a draw to 10:01:30.000
[09:58:00.000] The start time for the competitor(3) was set by a draw to 10:03:00.000
[09:59:30.000] The start time for the competitor(4) was set by a draw to 10:04:30.000
[09:59:45.000] The competitor(1) is on the start line
[10:00:01.744] The competitor(1) has started
[10:01:00.000] The start time for the competitor(5) was set by a draw to 10:06:00.000
[10:01:09.000] The competitor(2) is on the start line
[10:01:31.503] The competitor(2) has started
[10:02:36.000] The competitor(3) is on the start line
[10:03:00.887] The competitor(3) has started
[10:04:08.000] The competitor(4) is on the start line
[10:04:31.278] The competitor(4) has started
[10:05:42.000] The competitor(5) is on the start line
[10:06:00.331] The competitor(5) has started
[10:08:49.289] The competitor(1) is on the firing range(1)
[10:08:50.884] The target(1) has been hit by competitor(1)
[10:08:51.400] The target(2) has been hit by competitor(1)
[10:08:52.797] The target(5) has been hit by competitor(1)
[10:08:55.658] The competitor(1) left the firing range
[10:09:03.232] The competitor(1) entered the penalty laps
[10:10:22.273] The competitor(2) is on the firing range(1)
[10:10:23.804] The target(1) has been hit by competitor(2)
[10:10:25.036] The target(3) has been hit by competitor(2)
[10:10:25.449] The target(4) has been hit by competitor(2)
[10:10:26.002] The target(5) has been hit by competitor(2)
[10:10:29.125] The competitor(2) left the firing range
[10:10:38.142] The competitor(2) entered the penalty laps
[10:10:43.232] The competitor(1) left the penalty laps
[10:11:28.142] The competitor(2) left the penalty laps
[10:11:54.557] The competitor(3) is on the firing range(1)
[10:11:56.076] The target(1) has been hit by competitor(3)
[10:11:56.760] The target(2) has been hit by competitor(3)
[10:11:57.217] The target(3) has been hit by competitor(3)
[10:11:57.659] The target(4) has been hit by competitor(3)
[10:11:58.179] The target(5) has been hit by competitor(3)
[10:12:01.341] The competitor(3) left the firing range
[10:12:35.380] The competitor(1) ended the main lap
[10:13:27.246] The competitor(4) is on the firing range(1)
[10:13:29.773] The target(3) has been hit by competitor(4)
[10:13:30.443] The target(4) has been hit by competitor(4)
[10:13:30.836] The target(5) has been hit by competitor(4)
[10:13:33.970] The competitor(4) left the firing range
[10:13:43.912] The competitor(4) entered the penalty laps
[10:14:09.746] The competitor(2) ended the main lap
[10:15:20.988] The competitor(5) is on the firing range(1)
[10:15:22.758] The target(1) has been hit by competitor(5)
[10:15:23.083] The target(2) has been hit by competitor(5)
[10:15:23.682] The target(3) has been hit by competitor(5)
[10:15:23.912] The competitor(4) left the penalty laps
[10:15:27.197] The competitor(5) left the firing range
[10:15:31.757] The competitor(5) entered the penalty laps
[10:15:43.273] The competitor(3) ended the main lap
[10:17:11.757] The competitor(5) left the penalty laps
[10:17:16.947] The competitor(4) ended the main lap
[10:19:21.270] The competitor(5) ended the main lap
[10:21:34.847] The competitor(1) is on the firing range(2)
[10:21:36.495] The target(1) has been hit by competitor(1)
[10:21:36.920] The target(2) has been hit by competitor(1)
[10:21:37.626] The target(3) has been hit by competitor(1)
[10:21:38.628] The target(5) has been hit by competitor(1)
[10:21:41.449] The competitor(1) left the firing range
[10:21:50.476] The competitor(1) entered the penalty laps
[10:22:40.476] The competitor(1) left the penalty laps
[10:23:00.773] The competitor(2) is on the firing range(2)
[10:23:02.498] The target(1) has been hit by competitor(2)
[10:23:02.841] The target(2) has been hit by competitor(2)
[10:23:03.453] The target(3) has been hit by competitor(2)
[10:23:04.051] The target(4) has been hit by competitor(2)
[10:23:07.554] The competitor(2) left the firing range
[10:23:10.987] The competitor(2) entered the penalty laps
[10:24:00.987] The competitor(2) left the penalty laps
[10:24:43.323] The competitor(3) is on the firing range(2)
[10:24:44.954] The target(1) has been hit by competitor(3)
[10:24:45.508] The target(2) has been hit by competitor(3)
[10:24:45.923] The target(3) has been hit by competitor(3)
[10:24:46.559] The target(4) has been hit by competitor(3)
[10:24:46.958] The target(5) has been hit by competitor(3)
[10:24:49.905] The competitor(3) left the firing range
[10:25:26.047] The competitor(1) ended the main lap
[10:26:36.573] The competitor(4) is on the firing range(2)
[10:26:38.368] The target(1) has been hit by competitor(4)
[10:26:38.786] The target(2) has been hit by competitor(4)
[10:26:39.113] The target(3) has been hit by competitor(4)
[10:26:39.629] The target(4) has been hit by competitor(4)
[10:26:40.238] The target(5) has been hit by competitor(4)
[10:26:43.208] The competitor(4) left the firing range
[10:26:48.356] The competitor(2) ended the main lap
[10:28:28.112] The competitor(5) is on the firing range(2)
[10:28:29.629] The target(1) has been hit by competitor(5)
[10:28:30.408] The target(2) has been hit by competitor(5)
[10:28:30.769] The target(3) has been hit by competitor(5)
[10:28:31.882] The target(5) has been hit by competitor(5)
[10:28:34.274] The competitor(5) left the firing range
[10:28:34.773] The competitor(3) ended the main lap
[10:28:38.151] The competitor(5) entered the penalty laps
[10:29:28.151] The competitor(5) left the penalty laps
[10:30:36.413] The competitor(4) ended the main lap
[10:32:22.472] The competitor(5) ended the main lap
```
`result.txt` (файл результатов)
```
[00:25:18.356] 2 [ {00:12:38.243, 4.617}, {00:12:38.610, 4.617}] {{00:00:50.000, 3.000}, {00:00:50.000, 3.000}} 8/10 
[00:25:26.047] 1 [ {00:12:33.636, 4.648}, {00:12:50.667, 4.545}] {{00:01:40.000, 1.500}, {00:00:50.000, 3.000}} 7/10 
[00:25:34.773] 3 [ {00:12:42.386, 4.593}, {00:12:51.500, 4.540}] {} 10/10 
[00:26:06.413] 4 [ {00:12:45.669, 4.575}, {00:13:19.466, 4.380}] {{00:01:40.000, 1.500}} 8/10 
[00:26:22.472] 5 [ {00:13:20.939, 4.375}, {00:13:01.202, 4.481}] {{00:01:40.000, 1.500}, {00:00:50.000, 3.000}} 7/10 
```


# System prototype for biathlon competitions
The prototype must be able to work with a configuration file and a set of external events of a certain format.
Solution should contain golang (1.20 or newer) source file/files and unit tests (optional)

## Configuration (json)

- **Laps**        - Amount of laps for main distance
- **LapLen**      - Length of each main lap
- **PenaltyLen**  - Length of each penalty lap
- **FiringLines** - Number of firing lines per lap
- **Start**       - Planned start time for the first competitor
- **StartDelta**  - Planned interval between starts

## Events
All events are characterized by time and event identifier. Outgoing events are events created during program operation. Events related to the "incoming" category cannot be generated and are output in the same form as they were submitted in the input file.

- All events occur sequentially in time. (***Time of event N+1***) >= (***Time of event N***)
- Time format ***[HH:MM:SS.sss]***. Trailing zeros are required in input and output

#### Common format for events:
[***time***] **eventID** **competitorID** extraParams

```
Incoming events
EventID | extraParams | Comments
1       |             | The competitor registered
2       | startTime   | The start time was set by a draw
3       |             | The competitor is on the start line
4       |             | The competitor has started
5       | firingRange | The competitor is on the firing range
6       | target      | The target has been hit
7       |             | The competitor left the firing range
8       |             | The competitor entered the penalty laps
9       |             | The competitor left the penalty laps
10      |             | The competitor ended the main lap
11      | comment     | The competitor can`t continue
```
An competitor is disqualified if he/she does not start during his/her start interval. This marked as **NotStarted** in final report.
If the competitor can`t continue it should be marked in final report as **NotFinished**

```
Outgoing events
EventID | extraParams | Comments
32      |             | The competitor is disqualified
33      |             | The competitor has finished
```

## Final report
The final report should contain the list of all registered competitors
sorted by ascending time.
- Total time includes the difference between scheduled and actual start time or **NotStarted**/**NotFinished** marks
- Time taken to complete each lap
- Average speed for each lap [m/s]
- Time taken to complete penalty laps
- Average speed over penalty laps [m/s]
- Number of hits/number of shots

Examples:

`Config.conf`
```json
{
    "laps" : 2,
    "lapLen": 3651,
    "penaltyLen": 50,
    "firingLines": 1,
    "start": "09:30:00",
    "startDelta": "00:00:30"
}
```

`IncomingEvents`

```
[09:05:59.867] 1 1
[09:15:00.841] 2 1 09:30:00.000
[09:29:45.734] 3 1
[09:30:01.005] 4 1
[09:49:31.659] 5 1 1
[09:49:33.123] 6 1 1
[09:49:34.650] 6 1 2
[09:49:35.937] 6 1 4
[09:49:37.364] 6 1 5
[09:49:38.339] 7 1
[09:49:55.915] 8 1
[09:51:48.391] 9 1
[09:59:03.872] 10 1
[09:59:03.872] 11 1 Lost in the forest

```

`Output log`
```
[09:05:59.867] The competitor(1) registered
[09:15:00.841] The start time for the competitor(1) was set by a draw to 09:30:00.000
[09:29:45.734] The competitor(1) is on the start line
[09:30:01.005] The competitor(1) has started
[09:49:31.659] The competitor(1) is on the firing range(1)
[09:49:33.123] The target(1) has been hit by competitor(1)
[09:49:34.650] The target(2) has been hit by competitor(1)
[09:49:35.937] The target(4) has been hit by competitor(1)
[09:49:37.364] The target(5) has been hit by competitor(1)
[09:49:38.339] The competitor(1) left the firing range
[09:49:55.915] The competitor(1) entered the penalty laps
[09:51:48.391] The competitor(1) left the penalty laps
[09:59:03.872] The competitor(1) ended the main lap
[09:59:05.321] The competitor(1) can`t continue: Lost in the forest
```

`Resulting table`
```
[NotFinished] 1 [{00:29:03.872, 2.093}, {,}] {00:01:44.296, 0.481} 4/5