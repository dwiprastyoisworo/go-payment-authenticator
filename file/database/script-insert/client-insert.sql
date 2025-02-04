-- Client 1: Web Application
INSERT INTO public.clients (client_id, client_secret, name, redirect_uris)
VALUES (
           'web-app-123',
           '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', -- bcrypt('web-secret-456')
           'Web Application',
           ARRAY['http://localhost:3000/callback','https://webapp.com/callback']
       );

-- Client 2: Mobile App
INSERT INTO public.clients (client_id, client_secret, name, redirect_uris)
VALUES (
           'mobile-app-789',
           '$2a$10$sflkjsdlfkjlskdjflskjdf.4k1jdf8s7df4s8df7s8d4f7s8d', -- bcrypt('mobile-secret-abc')
           'Mobile Application',
           ARRAY['myapp://callback']
       );

-- Client 3: Internal Service (disabled)
INSERT INTO public.clients (client_id, client_secret, name, redirect_uris, enabled)
VALUES (
           'internal-service-101',
           '$2a$10$lkjflskdjflskdjflskdjf.5dkfj9d8fj9d8fj9d8fj9d8fj', -- bcrypt('internal-secret-202')
           'Internal Service',
           ARRAY['http://internal-service.local/callback'],
           false
       );

-- Client 4: Third-party Integration
INSERT INTO public.clients (client_id, client_secret, name, redirect_uris)
VALUES (
           'third-party-303',
           '$2a$10$lksjdflskdjflskdjflskd.s8d7f4s8d7f4s8d7f4s8d7f4', -- bcrypt('thirdparty-secret-404')
           'Third-party Integration',
           ARRAY['https://thirdparty.com/oauth/callback']
       );

-- Client 5: Test Environment
INSERT INTO public.clients (client_id, client_secret, name, redirect_uris)
VALUES (
           'test-env-505',
           '$2a$10$sldkfjslkdjflskdjflskdj.s7d8f4s7d8f4s7d8f4s7d8f4', -- bcrypt('test-secret-606')
           'Test Environment',
           ARRAY['http://test.local/callback','http://staging.test.local/callback']
       );