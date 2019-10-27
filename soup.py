#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Created on Sun Oct 27 03:10:49 2019

@author: alexey
"""

from bs4 import BeautifulSoup
import requests as req
#
##with open("https://amdm.ru/akkordi/yorsh/171254/iskri_novoy_zari/", "r") as f:
#resp = req.get('https://amdm.ru/akkordi/viktor_coi/5572/kamchatka/')
#
#soup = BeautifulSoup(resp.text, 'lxml')
#temp = soup.find("pre").contents[0::]
#print(soup.find("pre"))

def get_chords(search_req):
    resp = req.get('https://amdm.ru/search/?q='+search_req)
    #resp = req.get('https://amdm.ru/search/?q=43453ret')
    #soup.find_all("td", class_="sortcell")
    #print(resp)
    soup = BeautifulSoup(resp.text, 'lxml')
    table=soup.find('table')
    list_songs = []
    even = False
    for a in table.findAll('a', {'class':'artist'}, href=True):
        if even:
            list_songs.append('https:'+a['href'])
        even = not even
    resp = req.get(list_songs[0])
    soup = BeautifulSoup(resp.text, 'lxml')
    #output = str(soup.find("pre"))[28:-6:]
    #output = soup.find("pre").contents[1]
    output = soup.find("pre")
    return output
