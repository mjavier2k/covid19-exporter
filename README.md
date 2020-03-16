# Covid19-exporter
covid-19 prometheus exporter

# Requirements

1) Sign up on https://rapidapi.com/KishCom/api/covid-19-coronavirus-statistics 
2) Get you API Key
3) Create .env file

```
export RAPIDAPI_URL="https://covid-19-coronavirus-statistics.p.rapidapi.com/v1/stats"
export RAPIDAPI_HOST="covid-19-coronavirus-statistics.p.rapidapi.com"
export RAPIDAPI_KEY="XXXXXXXX"
```


# Example Metrics

```
covid19_deaths{country="Canada",province="Alberta"} 0
covid19_deaths{country="Canada",province="British Columbia"} 4
covid19_deaths{country="Canada",province="Manitoba"} 0
covid19_deaths{country="Canada",province="New Brunswick"} 0
covid19_deaths{country="Canada",province="Newfoundland and Labrador"} 0
covid19_deaths{country="Canada",province="Ontario"} 4
covid19_deaths{country="Canada",province="Prince Edward Island"} 0
covid19_deaths{country="Canada",province="Quebec"} 0
covid19_deaths{country="Canada",province="Saskatchewan"} 0
covid19_infected{country="Canada",province="Alberta"} 39
covid19_infected{country="Canada",province="British Columbia"} 73
covid19_infected{country="Canada",province="Manitoba"} 4
covid19_infected{country="Canada",province="New Brunswick"} 2
covid19_infected{country="Canada",province="Newfoundland and Labrador"} 1
covid19_infected{country="Canada",province="Ontario"} 104
covid19_infected{country="Canada",province="Prince Edward Island"} 1
covid19_infected{country="Canada",province="Quebec"} 24
covid19_infected{country="Canada",province="Saskatchewan"} 2
covid19_recovered{country="Canada",province="Alberta"} 0
covid19_recovered{country="Canada",province="British Columbia"} 4
covid19_recovered{country="Canada",province="Manitoba"} 0
covid19_recovered{country="Canada",province="New Brunswick"} 0
covid19_recovered{country="Canada",province="Newfoundland and Labrador"} 0
covid19_recovered{country="Canada",province="Ontario"} 4
covid19_recovered{country="Canada",province="Prince Edward Island"} 0
covid19_recovered{country="Canada",province="Quebec"} 0
covid19_recovered{country="Canada",province="Saskatchewan"} 0
```
