import requests
import json


header = {'TRN-Api-Key':'e51f2267-2f28-418f-8033-73a42b5f34c3'}
api = requests.get('https://api.fortnitetracker.com/v1/profile/pc/LopDropFlop', headers=header)
api = json.loads(api.text)
#print(api)


def main():

    while True:
        usr_input = input('''Type 'data' to see the hardcode API data, use 'api' to update API data from the API link or just type 'end' to break.

''')
        
        if usr_input == 'data':
            print ('''{'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'platformId': 3, 'platformName': 'pc', 'platformNameLong': 'PC', 'epicUserHandle': 'LopDropFlop', 'stats': {'p2': {'trnRating': {'label': 'TRN Rating', 'field': 'TRNRating', 'category': 'Rating', 'valueInt': 651, 'value': '651', 'rank': 232195, 'percentile': 71.0, 'displayValue': '651'}, 'score': {'label': 'Score', 'field': 'Score', 'category': 'General', 'valueInt': 45576, 'value': '45576', 'rank': 290164, 'percentile': 24.0, 'displayValue': '45,576'}, 'top1': {'label': 'Wins', 'field': 'Top1', 'category': 'Tops', 'valueInt': 12, 'value': '12', 'rank': 157019, 'percentile': 21.0, 'displayValue': '12'}, 'top3': {'label': 'Top 3', 'field': 'Top3', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'top5': {'label': 'Top 5', 'field': 'Top5', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'top6': {'label': 'Top 6', 'field': 'Top6', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'top10': {'label': 'Top 10', 'field': 'Top10', 'category': 'Tops', 'valueInt': 57, 'value': '57', 'rank': 207965, 'percentile': 4.2, 'displayValue': '57'}, 'top12': {'label': 'Top 12', 'field': 'Top12', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'top25': {'label': 'Top 25', 'field': 'Top25', 'category': 'Tops', 'valueInt': 102, 'value': '102', 'rank': 273333, 'percentile': 6.0, 'displayValue': '102'}, 'kd': {'label': 'K/d', 'field': 'KD', 'category': 'General', 'valueDec': 2.24, 'value': '2.24', 'rank': 194149, 'percentile': 11.0, 'displayValue': '2.24'}, 'winRatio': {'label': 'Win %', 'field': 'WinRatio', 'category': 'General', 'valueDec': 3.5, 'value': '3.5', 'rank': 268443, 'percentile': 30.0, 'displayValue': '3.50'}, 'matches': {'label': 'Matches', 'field': 'Matches', 'category': 'General', 'valueInt': 342, 'value': '342', 'rank': 285856, 'percentile': 26.0, 'displayValue': '342'}, 'kills': {'label': 'Kills', 'field': 'Kills', 'category': 'General', 'valueInt': 739, 'value': '739', 'rank': 157942, 'percentile': 11.0, 'displayValue': '739'}, 'minutesPlayed': {'label': 'Time Played', 'field': 'MinutesPlayed', 'category': 'General', 'valueInt': 2341, 'value': '2341', 'rank': 306963, 'percentile': 24.0, 'displayValue': '1d 15h 1m '}, 'kpm': {'label': 'Kills Per Min', 'field': 'KPM', 'category': 'General', 'valueDec': 0.32, 'value': '0.32', 'rank': 175958, 'percentile': 7.0, 'displayValue': '0.32'}, 'kpg': {'label': 'Kills Per Match', 'field': 'KPG', 'category': 'General', 'valueDec': 2.16, 'value': '2.16', 'rank': 191237, 'percentile': 10.0, 'displayValue': '2.16'}, 'avgTimePlayed': {'label': 'Avg Match Time', 'field': 'AvgTimePlayed', 'category': 'General', 'valueDec': 410.7, 'value': '410.7', 'rank': 881918, 'percentile': 64.0, 'displayValue': '6m 50s'}, 'scorePerMatch': {'label': 'Score per
        Match', 'field': 'ScorePerMatch', 'category': 'General', 'valueDec': 133.26, 'value': '133.26', 'rank': 759137, 'percentile': 55.0, 'displayValue': '133.26'}, 'scorePerMin': {'label': 'Score per Minute', 'field': 'ScorePerMin', 'category': 'General', 'valueDec': 19.47, 'value': '19.47', 'rank': 853633, 'percentile': 53.0, 'displayValue': '19.47'}}, 'p10': {'trnRating': {'label': 'TRN Rating', 'field': 'TRNRating', 'category': 'Rating', 'valueInt': 905, 'value': '905', 'rank': 202449, 'percentile': 66.0, 'displayValue': '905'}, 'score': {'label': 'Score', 'field': 'Score', 'category': 'General', 'valueInt': 40410, 'value': '40410', 'rank': 575009, 'percentile': 27.0, 'displayValue': '40,410'}, 'top1': {'label': 'Wins', 'field': 'Top1', 'category': 'Tops',
        'valueInt': 8, 'value': '8', 'rank': 426496, 'percentile': 27.0, 'displayValue': '8'}, 'top3': {'label': 'Top 3', 'field': 'Top3', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'top5': {'label': 'Top 5', 'field': 'Top5', 'category': 'Tops', 'valueInt': 24, 'value': '24', 'rank': 615687, 'percentile': 11.0, 'displayValue': '24'}, 'top6': {'label': 'Top 6', 'field': 'Top6', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'top10': {'label': 'Top 10', 'field': 'Top10', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'top12': {'label': 'Top 12', 'field': 'Top12', 'category': 'Tops', 'valueInt': 54, 'value': '54', 'rank': 698136, 'percentile': 0.0, 'displayValue': '54'}, 'top25': {'label': 'Top 25', 'field': 'Top25', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'kd': {'label': 'K/d', 'field': 'KD', 'category': 'General', 'valueDec': 2.38, 'value': '2.38', 'rank': 132653, 'percentile': 10.0, 'displayValue': '2.38'}, 'winRatio': {'label': 'Win %', 'field': 'WinRatio', 'category': 'General', 'valueDec': 4.1, 'value': '4.1', 'rank': 292809, 'percentile': 26.0, 'displayValue':
        '4.10'}, 'matches': {'label': 'Matches', 'field': 'Matches', 'category': 'General', 'valueInt': 193, 'value': '193', 'rank': 761051, 'percentile': 37.0,
        'displayValue': '193'}, 'kills': {'label': 'Kills', 'field': 'Kills', 'category': 'General', 'valueInt': 440, 'value': '440', 'rank': 433592, 'percentile': 21.0, 'displayValue': '440'}, 'minutesPlayed': {'label': 'Time Played', 'field': 'MinutesPlayed', 'category': 'General', 'valueInt': 1485, 'value': '1485', 'rank': 738565, 'percentile': 37.0, 'displayValue': '1d 45m '}, 'kpm': {'label': 'Kills Per Min', 'field': 'KPM', 'category': 'General', 'valueDec': 0.3, 'value': '0.3', 'rank': 110258, 'percentile': 8.0, 'displayValue': '0.30'}, 'kpg': {'label': 'Kills Per Match', 'field': 'KPG', 'category': 'General', 'valueDec': 2.28, 'value': '2.28', 'rank': 122202, 'percentile': 9.0, 'displayValue': '2.28'}, 'avgTimePlayed': {'label': 'Avg Match Time', 'field': 'AvgTimePlayed', 'category': 'General', 'valueDec': 461.66, 'value': '461.66', 'rank': 630705, 'percentile': 45.0, 'displayValue': '7m 41s'}, 'scorePerMatch': {'label': 'Score per Match', 'field': 'ScorePerMatch', 'category': 'General', 'valueDec': 209.38, 'value': '209.38', 'rank': 246263, 'percentile': 17.0, 'displayValue': '209.38'}, 'scorePerMin': {'label': 'Score per Minute', 'field': 'ScorePerMin', 'category': 'General', 'valueDec': 27.21, 'value': '27.21', 'rank': 215947, 'percentile': 12.0, 'displayValue': '27.21'}}, 'p9': {'trnRating': {'label': 'TRN Rating', 'field': 'TRNRating', 'category': 'Rating', 'valueInt': 1124, 'value': '1124', 'rank': 177671, 'percentile': 58.0, 'displayValue': '1,124'}, 'score': {'label': 'Score', 'field': 'Score', 'category': 'General', 'valueInt': 90245, 'value': '90245', 'rank': 324908, 'percentile': 10.0, 'displayValue': '90,245'}, 'top1': {'label': 'Wins', 'field': 'Top1', 'category': 'Tops', 'valueInt': 28, 'value': '28', 'rank': 307124, 'percentile': 10.0, 'displayValue': '28'}, 'top3': {'label': 'Top 3', 'field': 'Top3', 'category': 'Tops', 'valueInt': 56, 'value': '56', 'rank': 389154, 'percentile': 7.0, 'displayValue': '56'}, 'top5': {'label': 'Top 5', 'field': 'Top5', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'top6': {'label': 'Top 6', 'field': 'Top6', 'category': 'Tops', 'valueInt': 89, 'value': '89', 'rank': 420257, 'percentile': 8.0, 'displayValue': '89'}, 'top10': {'label': 'Top 10', 'field': 'Top10', 'category':
        'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'top12': {'label': 'Top 12', 'field': 'Top12', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1, 'displayValue': '0'}, 'top25': {'label': 'Top 25', 'field': 'Top25', 'category': 'Tops', 'valueInt': 0, 'value': '0', 'rank': 1,
        'displayValue': '0'}, 'kd': {'label': 'K/d', 'field': 'KD', 'category': 'General', 'valueDec': 2.68, 'value': '2.68', 'rank': 101730, 'percentile': 7.0,
        'displayValue': '2.68'}, 'winRatio': {'label': 'Win %', 'field': 'WinRatio', 'category': 'General', 'valueDec': 8.2, 'value': '8.2', 'rank': 278058, 'percentile': 11.0, 'displayValue': '8.20'}, 'matches': {'label': 'Matches', 'field': 'Matches', 'category': 'General', 'valueInt': 342, 'value': '342', 'rank': 505662, 'percentile': 26.0, 'displayValue': '342'}, 'kills': {'label': 'Kills', 'field': 'Kills', 'category': 'General', 'valueInt': 840, 'value': '840', 'rank': 231668, 'percentile': 10.0, 'displayValue': '840'}, 'minutesPlayed': {'label': 'Time Played', 'field': 'MinutesPlayed', 'category': 'General', 'valueInt': 2798, 'value': '2798', 'rank': 510686, 'percentile': 19.0, 'displayValue': '1d 22h 38m '}, 'kpm': {'label': 'Kills Per Min', 'field': 'KPM', 'category': 'General', 'valueDec': 0.3, 'value': '0.3', 'rank': 79993, 'percentile': 8.0, 'displayValue': '0.30'}, 'kpg': {'label': 'Kills Per Match', 'field': 'KPG', 'category': 'General', 'valueDec': 2.46, 'value': '2.46', 'rank': 90107, 'percentile': 7.0, 'displayValue': '2.46'}, 'avgTimePlayed': {'label': 'Avg Match Time', 'field': 'AvgTimePlayed', 'category': 'General', 'valueDec': 490.88, 'value': '490.88', 'rank': 763599, 'percentile': 34.0, 'displayValue': '8m 10s'}, 'scorePerMatch': {'label': 'Score per Match', 'field': 'ScorePerMatch', 'category': 'General', 'valueDec': 263.87, 'value': '263.87', 'rank': 179563, 'percentile': 6.0, 'displayValue': '263.87'}, 'scorePerMin': {'label': 'Score per Minute', 'field': 'ScorePerMin', 'category': 'General', 'valueDec': 32.25, 'value': '32.25', 'rank': 123841, 'percentile': 4.0, 'displayValue': '32.25'}}}, 'lifeTimeStats': [{'key': 'Top 3', 'value': '57'}, {'key': 'Top 5s', 'value': '24'}, {'key': 'Top 3s', 'value': '56'}, {'key': 'Top 6s', 'value': '89'}, {'key': 'Top 12s', 'value': '54'}, {'key': 'Top 25s', 'value': '102'}, {'key': 'Score', 'value': '176,231'}, {'key': 'Matches Played', 'value': '877'}, {'key': 'Wins', 'value': '48'}, {'key': 'Win%',
        'value': '6%'}, {'key': 'Kills', 'value': '2019'}, {'key': 'K/d', 'value': '2.44'}, {'key': 'Kills Per Min', 'value': '0.3'}, {'key': 'Time Played', 'value': '4d 14h 24m '}, {'key': 'Avg Survival Time', 'value': '7m 33s'}], 'recentMatches': [{'id': 25763742, 'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'playlist': 'p2', 'kills': 6, 'minutesPlayed': 10, 'top1': 0, 'top5': 0, 'top6': 0, 'top10': 0, 'top12': 0, 'top25': 0, 'matches': 2, 'top3': 0, 'dateCollected': '2018-02-16T02:09:19.08', 'score': 336, 'platform': 3, 'trnRating': 651.7, 'trnRatingChange': 32.858945625}, {'id': 25755259, 'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'playlist': 'p10', 'kills': 2, 'minutesPlayed': 15, 'top1': 0, 'top5': 0, 'top6': 0, 'top10': 0, 'top12': 0, 'top25': 0, 'matches': 5, 'top3': 0, 'dateCollected': '2018-02-16T01:54:09.25', 'score': 314, 'platform': 3, 'trnRating': 905.3, 'trnRatingChange': -99.87023081152343}, {'id': 25755260, 'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'playlist': 'p2', 'kills': 11, 'minutesPlayed': 18, 'top1': 0, 'top5': 0, 'top6': 0, 'top10': 0, 'top12': 0, 'top25': 0, 'matches': 5, 'top3': 0, 'dateCollected': '2018-02-16T01:54:09.25', 'score': 567, 'platform': 3, 'trnRating': 585.7, 'trnRatingChange': 33.357234335083476}, {'id': 25677939, 'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'playlist': 'p9', 'kills': 68, 'minutesPlayed': 272, 'top1': 2, 'top5': 0, 'top6': 7, 'top10': 0, 'top12': 0, 'top25': 0, 'matches': 35, 'top3': 6, 'dateCollected': '2018-02-15T23:45:43.273', 'score': 9207, 'platform': 3, 'trnRating': 1124.4, 'trnRatingChange': -9.664762341108691}, {'id': 25677940, 'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'playlist': 'p2', 'kills': 13, 'minutesPlayed': 43, 'top1': 0, 'top5': 0, 'top6': 0, 'top10': 0, 'top12': 0, 'top25': 1, 'matches': 9, 'top3': 0, 'dateCollected': '2018-02-15T23:45:43.273', 'score': 1097, 'platform': 3, 'trnRating': 416.4, 'trnRatingChange': -111.79589513249545}, {'id': 25677938, 'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'playlist': 'p10', 'kills': 119, 'minutesPlayed': 395, 'top1': 1, 'top5': 4, 'top6': 0, 'top10': 0, 'top12': 15, 'top25': 0, 'matches': 50, 'top3': 0, 'dateCollected': '2018-02-15T23:45:43.257', 'score': 11569, 'platform': 3, 'trnRating': 1380.9, 'trnRatingChange': 2.7163430567520193}, {'id': 22549789, 'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'playlist': 'p9', 'kills': 12, 'minutesPlayed': 34, 'top1': 1, 'top5': 0, 'top6': 1, 'top10': 0, 'top12': 0, 'top25': 0, 'matches': 3, 'top3': 1, 'dateCollected': '2018-02-12T01:26:42.17', 'score': 1353, 'platform': 3, 'trnRating': 1316.7, 'trnRatingChange': 27.833669999999998}, {'id': 22504047, 'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'playlist': 'p9', 'kills': 1, 'minutesPlayed': 3, 'top1': 0, 'top5': 0, 'top6': 0, 'top10': 0, 'top12': 0, 'top25': 0, 'matches': 1, 'top3': 0, 'dateCollected': '2018-02-12T00:17:06.907', 'score': 117, 'platform': 3, 'trnRating': 1105.0, 'trnRatingChange': -95.0}, {'id': 22501016, 'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'playlist': 'p2', 'kills': 709, 'minutesPlayed': 2270, 'top1': 12, 'top5': 0, 'top6': 0, 'top10': 57, 'top12': 0, 'top25': 101, 'matches': 326, 'top3': 0, 'dateCollected': '2018-02-12T00:12:20.393', 'score': 43576, 'platform': 3}, {'id': 22501012, 'accountId': '1d3cf41d-1627-47ca-a947-e1d31305d2b4', 'playlist': 'p10', 'kills': 319, 'minutesPlayed': 1075, 'top1': 7, 'top5': 20, 'top6': 0, 'top10': 0, 'top12': 39, 'top25': 0, 'matches': 138, 'top3': 0, 'dateCollected': '2018-02-12T00:12:20.377', 'score': 28527, 'platform': 3}]} ''')
        elif usr_input == 'api':
            print(api)
        elif usr_input == 'end':
            break

main()