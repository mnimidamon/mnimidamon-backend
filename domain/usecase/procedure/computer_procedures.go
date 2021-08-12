package procedure

import (
	"errors"
	. "mnimidamonbackend/domain"
	"mnimidamonbackend/domain/model"
	"mnimidamonbackend/domain/repository"
)

func DeleteComputer(c *model.Computer, gcr repository.GroupComputerRepository, cbr repository.ComputerBackupRepository, cr repository.ComputerRepository, br repository.BackupRepository) error {
	// Get all group computers.
	gcs, err := gcr.FindAllOfComputer(c.ID)
	if err != nil {
		return ToDomainError(err)
	}

	// For each group computer of the computer.
	for _, gc := range gcs {
		cbs, err := cbr.FindAllOfGroupComputer(gc.ID)
		if err != nil {
			return ToDomainError(err)
		}

		backupIDS := []uint{}
		// Delete each computer backup.
		for _, cb := range cbs {
			backupIDS = append(backupIDS, cb.BackupID)
			if err := cbr.Delete(cb.GroupComputerID, cb.BackupID); err != nil {
				return ToDomainError(err)
			}
		}

		// Add the IDs of the backups inside the group.
		groupBackups, err := br.FindAll(gc.GroupID)
		if err != nil {
			return ToDomainError(err)
		}

		for _, b := range groupBackups {
			if b.OwnerID == c.OwnerID {
				backupIDS = append(backupIDS, b.ID)
			}
		}

		// Delete the group computer.
		if err := gcr.Delete(gc.GroupID, gc.ComputerID); err != nil {
			return ToDomainError(err)
		}

		// Delete the backup if none have it, or the group does not have any group computer of the owner inside.
		for _, bID := range backupIDS {
			b, err := br.FindById(bID)
			if err != nil {
				// Might have duplicates here, so lets skip the ones that were deleted.
				if errors.Is(err, repository.ErrNotFound) {
					continue
				}
				return ToDomainError(err)
			}

			isStillComputerMember := false
			// If none is member of group.
			computers, err := cr.FindAll(c.OwnerID)
			if err != nil {
				return ToDomainError(err)
			}

			for _, cm := range computers {
				groupComputers, err := gcr.FindAllOfComputer(cm.ID)
				if err != nil {
					return ToDomainError(err)
				}

				// Is still a member with any of these computers?
				for _, gc := range groupComputers {
					// At least one group computer where the backup is located.
					if gc.GroupID == b.GroupID {
						isStillComputerMember = true
						break
					}
				}
			}

			// If the backup does not have any backupees left, just delete it.
			backupees, err := cbr.FindAllOfBackup(bID)
			if err != nil {
				return err
			}

			// Delete the backup if hes not a memeber anymore or no one has this backup.
			if !isStillComputerMember || len(backupees) == 0 {
				if err := br.Delete(bID); err != nil {
					return err
				}
			}
		}
	}

	// Finally delete the computer
	if err := cr.Delete(c.ID); err != nil {
		return ToDomainError(err)
	}

	return nil
}
