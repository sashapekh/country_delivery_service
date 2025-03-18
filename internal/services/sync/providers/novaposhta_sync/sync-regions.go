package novaposhta_sync

func (h *SyncServiceHanlder) SyncRegions() error {

	err := SyncRegionsNovaposhta(
		h.novaposhta,
		h.RepoHandler,
	)

	if err != nil {
		h.logger.Error(err.Error())
		return err
	}

	return nil
}
