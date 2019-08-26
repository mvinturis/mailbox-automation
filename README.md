mailbox-automation
==================

mailbox-automation implements an automated runner to simulate user interaction with a specific mailbox (Hotmail, Yahoo, AOL). The program uses Chromedp GO package to develop activities for each runner implementation. In general the runner activities are comprised of:
- mark not spam emails from the Spam folder;
- open messages from Inbox folder;
- pin/categorize/flag messages from Inbox folder;
- move Inbox messages to Archive folder;

Usage example:

Start the runner with a specified behavior, also configuring the seed/mailbox to run and additional parameters if needed:

.. code-block:: go

	seed := &models.Seed{
		Email:        "test1@hotmail.com"
		Password:     "test",
		RecoveryCode: "",
		LocalEmail:   "", // Recovery Email
		ProxyIp:      "",
	}
	params := &models.TaskParams{
		Keyword: "keyword to search",
	}

	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
	defer cancel()

	runner := hotmail.NewRunner(seed, taskCtx)

	runner.Start("readMessages", params)
