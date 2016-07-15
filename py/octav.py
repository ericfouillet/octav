"""OCTAV Client Library"""
"""DO NOT EDIT: This file was generated from ../spec/v1/api.json on Fri Jul 15 10:46:35 2016"""

import json
import os
import requests

class Octav(object):
  def __init__(self, endpoint=None, key=None, secret=None, debug=False):
    if not endpoint:
      raise "endpoint is required"
    if not key:
      raise "key is required"
    if not secret:
      raise "secret is required"
    self.debug = debug
    self.endpoint = endpoint
    self.error = None
    self.key = key
    self.secret = secret
    self.session = requests.Session()
    self.session.mount('http://', requests.adapters.HTTPAdapter(max_retries=0))

  def extract_error(self, r):
    try:
      js = r.json()
      self.error = js["message"]
    except:
      self.error = r.status_code

  def last_error(self):
    return self.error

  def create_user (self, last_name=None, first_name=None, auth_via=None, avatar_url=None, tshirt_size=None, auth_user_id=None, email=None, nickname=None):
    payload = {}
    if auth_user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['auth_user_id'] = auth_user_id
    if auth_via is None:
            raise "property \"" + required + "\" must be provided"
    payload['auth_via'] = auth_via
    if nickname is None:
            raise "property \"" + required + "\" must be provided"
    payload['nickname'] = nickname
    if not avatar_url is None:
        payload['avatar_url'] = avatar_url
    if not email is None:
        payload['email'] = email
    if not first_name is None:
        payload['first_name'] = first_name
    if not last_name is None:
        payload['last_name'] = last_name
    if not tshirt_size is None:
        payload['tshirt_size'] = tshirt_size
    uri = self.endpoint + "/user/create"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def lookup_user (self, id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    uri = self.endpoint + "/user/lookup"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def lookup_user_by_auth_user_id (self, auth_user_id=None, auth_via=None):
    payload = {}
    if auth_user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['auth_user_id'] = auth_user_id
    if auth_via is None:
            raise "property \"" + required + "\" must be provided"
    payload['auth_via'] = auth_via
    uri = self.endpoint + "/user/lookup_by_auth_user_id"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def update_user (self, email=None, id=None, nickname=None, user_id=None, last_name=None, tshirt_size=None, first_name=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if not email is None:
        payload['email'] = email
    if not first_name is None:
        payload['first_name'] = first_name
    if not last_name is None:
        payload['last_name'] = last_name
    if not nickname is None:
        payload['nickname'] = nickname
    if not tshirt_size is None:
        payload['tshirt_size'] = tshirt_size
    uri = self.endpoint + "/user/update"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def delete_user (self, user_id=None, id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/user/delete"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def list_user (self, limit=None, since=None, lang=None):
    payload = {}
    if not lang is None:
        payload['lang'] = lang
    if not limit is None:
        payload['limit'] = limit
    if not since is None:
        payload['since'] = since
    uri = self.endpoint + "/user/list"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def create_venue (self, latitude=None, name=None, longitude=None, address=None, user_id=None):
    payload = {}
    if address is None:
            raise "property \"" + required + "\" must be provided"
    payload['address'] = address
    if name is None:
            raise "property \"" + required + "\" must be provided"
    payload['name'] = name
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if not latitude is None:
        payload['latitude'] = latitude
    if not longitude is None:
        payload['longitude'] = longitude
    uri = self.endpoint + "/venue/create"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def list_venue (self, limit=None, since=None, lang=None):
    payload = {}
    if not lang is None:
        payload['lang'] = lang
    if not limit is None:
        payload['limit'] = limit
    if not since is None:
        payload['since'] = since
    uri = self.endpoint + "/venue/list"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def lookup_venue (self, id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    uri = self.endpoint + "/venue/lookup"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def update_venue (self, id=None, user_id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/venue/update"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def delete_venue (self, user_id=None, id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/venue/delete"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def create_room (self, venue_id=None, name=None, capacity=None, user_id=None):
    payload = {}
    if name is None:
            raise "property \"" + required + "\" must be provided"
    payload['name'] = name
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if venue_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['venue_id'] = venue_id
    if not capacity is None:
        payload['capacity'] = capacity
    uri = self.endpoint + "/room/create"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def update_room (self, user_id=None, venue_id=None, id=None, name=None, capacity=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if not capacity is None:
        payload['capacity'] = capacity
    if not name is None:
        payload['name'] = name
    if not venue_id is None:
        payload['venue_id'] = venue_id
    uri = self.endpoint + "/room/update"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def lookup_room (self, id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    uri = self.endpoint + "/room/lookup"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def delete_room (self, user_id=None, id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/room/delete"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def list_room (self, limit=None, venue_id=None, lang=None):
    payload = {}
    if venue_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['venue_id'] = venue_id
    if not lang is None:
        payload['lang'] = lang
    if not limit is None:
        payload['limit'] = limit
    uri = self.endpoint + "/room/list"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def create_conference_series (self, user_id=None, slug=None):
    payload = {}
    if slug is None:
            raise "property \"" + required + "\" must be provided"
    payload['slug'] = slug
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/conference_series/create"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def list_conference_series (self, since=None, limit=None):
    payload = {}
    if not limit is None:
        payload['limit'] = limit
    if not since is None:
        payload['since'] = since
    uri = self.endpoint + "/conference_series/list"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def add_conference_series_admin (self, user_id=None, series_id=None, admin_id=None):
    payload = {}
    if admin_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['admin_id'] = admin_id
    if series_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['series_id'] = series_id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/conference_series/admin/add"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def create_conference (self, user_id=None, series_id=None, description=None, title=None, slug=None, sub_title=None):
    payload = {}
    if series_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['series_id'] = series_id
    if slug is None:
            raise "property \"" + required + "\" must be provided"
    payload['slug'] = slug
    if title is None:
            raise "property \"" + required + "\" must be provided"
    payload['title'] = title
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if not description is None:
        payload['description'] = description
    if not sub_title is None:
        payload['sub_title'] = sub_title
    uri = self.endpoint + "/conference/create"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def add_conference_dates (self, dates=None, conference_id=None, user_id=None):
    payload = {}
    if conference_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['conference_id'] = conference_id
    if dates is None:
            raise "property \"" + required + "\" must be provided"
    payload['dates'] = dates
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/conference/dates/add"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def delete_conference_dates (self, user_id=None, conference_id=None, dates=None):
    payload = {}
    if conference_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['conference_id'] = conference_id
    if dates is None:
            raise "property \"" + required + "\" must be provided"
    payload['dates'] = dates
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/conference/dates/delete"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def add_conference_admin (self, user_id=None, conference_id=None, admin_id=None):
    payload = {}
    if admin_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['admin_id'] = admin_id
    if conference_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['conference_id'] = conference_id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/conference/admin/add"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def delete_conference_admin (self, user_id=None, conference_id=None, admin_id=None):
    payload = {}
    if admin_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['admin_id'] = admin_id
    if conference_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['conference_id'] = conference_id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/conference/admin/delete"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def add_conference_venue (self, user_id=None, conference_id=None, venue_id=None):
    payload = {}
    if conference_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['conference_id'] = conference_id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if venue_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['venue_id'] = venue_id
    uri = self.endpoint + "/conference/venue/add"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def delete_conference_venue (self, user_id=None, conference_id=None, venue_id=None):
    payload = {}
    if conference_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['conference_id'] = conference_id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if venue_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['venue_id'] = venue_id
    uri = self.endpoint + "/conference/venue/delete"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def lookup_conference (self, id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    uri = self.endpoint + "/conference/lookup"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def lookup_conference_by_slug (self, slug=None, lang=None):
    payload = {}
    if slug is None:
            raise "property \"" + required + "\" must be provided"
    payload['slug'] = slug
    if not lang is None:
        payload['lang'] = lang
    uri = self.endpoint + "/conference/lookup_by_slug"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def list_conference (self, limit=None, since=None, lang=None, range_start=None, range_end=None):
    payload = {}
    if not lang is None:
        payload['lang'] = lang
    if not limit is None:
        payload['limit'] = limit
    if not range_end is None:
        payload['range_end'] = range_end
    if not range_start is None:
        payload['range_start'] = range_start
    if not since is None:
        payload['since'] = since
    uri = self.endpoint + "/conference/list"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def update_conference (self, user_id=None, description=None, starts_on=None, title=None, id=None, sub_title=None, slug=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if not description is None:
        payload['description'] = description
    if not slug is None:
        payload['slug'] = slug
    if not starts_on is None:
        payload['starts_on'] = starts_on
    if not sub_title is None:
        payload['sub_title'] = sub_title
    if not title is None:
        payload['title'] = title
    uri = self.endpoint + "/conference/update"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def delete_conference_series (self, id=None, user_id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/conference_series/delete"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def delete_conference (self, id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    uri = self.endpoint + "/conference/delete"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def create_session (self, user_id=None, slide_subtitles=None, memo=None, spoken_language=None, category=None, slide_language=None, duration=None, tags=None, video_url=None, video_permission=None, speaker_id=None, title=None, conference_id=None, material_level=None, photo_permission=None, abstract=None, slide_url=None):
    payload = {}
    if abstract is None:
            raise "property \"" + required + "\" must be provided"
    payload['abstract'] = abstract
    if conference_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['conference_id'] = conference_id
    if duration is None:
            raise "property \"" + required + "\" must be provided"
    payload['duration'] = duration
    if speaker_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['speaker_id'] = speaker_id
    if title is None:
            raise "property \"" + required + "\" must be provided"
    payload['title'] = title
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if not category is None:
        payload['category'] = category
    if not material_level is None:
        payload['material_level'] = material_level
    if not memo is None:
        payload['memo'] = memo
    if not photo_permission is None:
        payload['photo_permission'] = photo_permission
    if not slide_language is None:
        payload['slide_language'] = slide_language
    if not slide_subtitles is None:
        payload['slide_subtitles'] = slide_subtitles
    if not slide_url is None:
        payload['slide_url'] = slide_url
    if not spoken_language is None:
        payload['spoken_language'] = spoken_language
    if not tags is None:
        payload['tags'] = tags
    if not video_permission is None:
        payload['video_permission'] = video_permission
    if not video_url is None:
        payload['video_url'] = video_url
    uri = self.endpoint + "/session/create"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def lookup_session (self, id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    uri = self.endpoint + "/session/lookup"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def delete_session (self, id=None, user_id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/session/delete"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def update_session (self, conference_id=None, material_level=None, has_interpretation=None, photo_permission=None, abstract=None, slide_url=None, tags=None, video_url=None, title=None, speaker_id=None, video_permission=None, category=None, slide_language=None, sort_order=None, status=None, duration=None, id=None, user_id=None, slide_subtitles=None, memo=None, confirmed=None, spoken_language=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if not abstract is None:
        payload['abstract'] = abstract
    if not category is None:
        payload['category'] = category
    if not conference_id is None:
        payload['conference_id'] = conference_id
    if not confirmed is None:
        payload['confirmed'] = confirmed
    if not duration is None:
        payload['duration'] = duration
    if not has_interpretation is None:
        payload['has_interpretation'] = has_interpretation
    if not material_level is None:
        payload['material_level'] = material_level
    if not memo is None:
        payload['memo'] = memo
    if not photo_permission is None:
        payload['photo_permission'] = photo_permission
    if not slide_language is None:
        payload['slide_language'] = slide_language
    if not slide_subtitles is None:
        payload['slide_subtitles'] = slide_subtitles
    if not slide_url is None:
        payload['slide_url'] = slide_url
    if not sort_order is None:
        payload['sort_order'] = sort_order
    if not speaker_id is None:
        payload['speaker_id'] = speaker_id
    if not spoken_language is None:
        payload['spoken_language'] = spoken_language
    if not status is None:
        payload['status'] = status
    if not tags is None:
        payload['tags'] = tags
    if not title is None:
        payload['title'] = title
    if not video_permission is None:
        payload['video_permission'] = video_permission
    if not video_url is None:
        payload['video_url'] = video_url
    uri = self.endpoint + "/session/update"
    if self.debug:
        print("POST " + uri)
    res = self.session.post(uri, auth=(self.key, self.secret), json=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def list_session_by_conference (self, date=None, conference_id=None):
    payload = {}
    if conference_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['conference_id'] = conference_id
    if not date is None:
        payload['date'] = date
    uri = self.endpoint + "/schedule/list"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def create_question (self, session_id=None, user_id=None, body=None):
    payload = {}
    if body is None:
            raise "property \"" + required + "\" must be provided"
    payload['body'] = body
    if session_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['session_id'] = session_id
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    uri = self.endpoint + "/question/create"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def delete_question (self, id=None):
    payload = {}
    if id is None:
            raise "property \"" + required + "\" must be provided"
    payload['id'] = id
    uri = self.endpoint + "/question/delete"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True

  def list_question (self, session_id=None, since=None, limit=None):
    payload = {}
    if session_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['session_id'] = session_id
    if not limit is None:
        payload['limit'] = limit
    if not since is None:
        payload['since'] = since
    uri = self.endpoint + "/question/list"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return res.json()

  def create_session_survey_response (self, comment_improvement=None, user_id=None, comment_good=None, user_prior_knowledge=None, overall_rating=None, speaker_presentation=None, session_id=None, speaker_knowledge=None, material_quality=None):
    payload = {}
    if material_quality is None:
            raise "property \"" + required + "\" must be provided"
    payload['material_quality'] = material_quality
    if overall_rating is None:
            raise "property \"" + required + "\" must be provided"
    payload['overall_rating'] = overall_rating
    if session_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['session_id'] = session_id
    if speaker_knowledge is None:
            raise "property \"" + required + "\" must be provided"
    payload['speaker_knowledge'] = speaker_knowledge
    if speaker_presentation is None:
            raise "property \"" + required + "\" must be provided"
    payload['speaker_presentation'] = speaker_presentation
    if user_id is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_id'] = user_id
    if user_prior_knowledge is None:
            raise "property \"" + required + "\" must be provided"
    payload['user_prior_knowledge'] = user_prior_knowledge
    if not comment_good is None:
        payload['comment_good'] = comment_good
    if not comment_improvement is None:
        payload['comment_improvement'] = comment_improvement
    uri = self.endpoint + "/survey_session_response/create"
    if self.debug:
        print("GET " + uri)
    res = self.session.get(uri, params=payload)
    if res.status_code != 200:
        self.extract_error(res)
        return None
    return True
