
insert into step (id, title, description, priority, actions, templates, next_steps, render_type, rule) values
(0, 'How can I help you?', 'Ask PayMate to help you with all your money related issues.', 1, '', '', '1,2,3,4', 'TEXT', '');

insert into step (id, title, description, priority, actions, templates, next_steps, render_type, rule) values
(1, 'Get Money', 'Request money from your friends or family', 1, '', '', '', 'LIST_ELEMENT', ''),
(2, 'Send Money', 'Send money to your friends or family', 2, '', '', '', 'LIST_ELEMENT', 'BALANCE_CHECK'),
(3, 'Request for a vouch', 'Have low balance? Ask your friends or family to vouch for you to get money from Grab', 3, '', '', '', 'LIST_ELEMENT', 'BALANCE_CHECK_X'),
(4, 'Lend money', 'Request money from your friends or family', 4, '', '', '', 'LIST_ELEMENT', 'BALANCE_CHECK');