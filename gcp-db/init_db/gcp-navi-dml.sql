
insert into questions(id, question_id, question, category_id) VALUES (
    1,
    '1',
	'GCPでハードディスクを暗号化するための推奨される方法はどれですか？',
    '1'
);


insert into questions(id, question_id, question, category_id) VALUES (
    2,
    '2',
	'Google Cloud SQLでサポートされているプログラミング言語はどれですか？',
    '1'
);

insert into answers (id,question_id,correct_answer,answer1,answer2,answer3,answer4,answer5,answer6,answer7,answer8,answer9,answer10)
VALUES (
        1,
		1,
        4,
		'Key Management Service (KMS)を使用する',
		'Google Cloud Storageの暗号化を使用する',
		'サードパーティのディスク暗号化製品を使用する',
		'Identity and Access Management（IAM）を使用する',
        '',
        '',
        '',
        '',
        '',
        ''
);

insert into answers (id,question_id,correct_answer,answer1,answer2,answer3,answer4,answer5,answer6,answer7,answer8,answer9,answer10)
VALUES (
        2,
		2,
        2,
		'Key Management Service (KMS)を使用しない',
		'Google Cloud Storageの暗号化を使用しない',
		'サードパーティのディスク暗号化製品を使用しない',
		'Identity and Access Management（IAM）を使用しない',
        '',
        '',
        '',
        '',
        '',
        ''
);

insert into category (id, category_id, category) values (1,1,'セキュリティ');
