package procedure

import (
	. "mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
)

func DeleteComputer(c *model.Computer, gcr repository.GroupComputerRepository, cbr repository.ComputerBackupRepository, cr repository.ComputerRepository) error {
	// Get all group computers.
	gcs, err := gcr.FindAllOfComputer(c.ID)
	if err != nil {
		return ToDomainError(err)
	}

	// For each computer get the backups.
	for _, gc := range gcs {
		bcs, err := cbr.FindAllOfGroupComputer(gc.ID)
		if err != nil {
			return ToDomainError(err)
		}

		// Delete each.
		for _, bc := range bcs {
			if err := cbr.Delete(bc.GroupComputerID, bc.BackupID); err != nil {
				return ToDomainError(err)
			}
		}

		// Delete the each group computer.
		if err := gcr.Delete(gc.GroupID, gc.ComputerID); err != nil {
			return ToDomainError(err)
		}
	}

	// Finally delete the computer
	if err := cr.Delete(c.ID); err != nil {
		return ToDomainError(err)
	}

	return nil
}
