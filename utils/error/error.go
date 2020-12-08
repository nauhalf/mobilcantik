package errorutils

const StatusZeroAffectedRows = "Zero affected rows"

const (
	ErrorMySQLDeleteConstraintFK       = 1451
	ErrorMySQLInsertUpdateConstraintFK = 1452
)

const (
	StatusErrorFillRequiredForm                 = "All required form must be filled"
	StatusErrorResourceNotExists                = "%s not exists"
	StatusErrorImageSize                        = "Image size must less than %s"
	StatusErrorImageType                        = "Image type must be PNG/JPG"
	StatusErrorInvalidYear                      = "Invalid year, must be between 1800 and current year"
	StatusErrorInvalidEmail                     = "Invalid email format"
	StatusErrorAlreadyInUsed                    = "%s is already in used, failed to delete it"
	StatusErrorSuccessfullyDeleted              = "%s is already in used, failed to delete it"
	StatusErrorSuccessfullyRetrieved            = "%s successfully retrieved"
	StatusErrorSuccessfullyCreated              = "%s successfully created"
	StatusErrorSuccessfullyUpdated              = "%s successfully updated"
	StatusErrorSuccessfullyConfirmed            = "%s successfully confirmed"
	StatusErrorIncorrectPassword                = "Incorrect password"
	StatusErrorAlreadyActive                    = "%s already actived"
	StatusErrorPaidToActivateAd                 = "You must paid the fee to activate your Ad"
	StatusErrorNotInWaitingFeeConfirmationState = "Your ad is not in waiting confirmation fee state"
	StatusErrorNotSame                          = "%s not same on the database"
	StatusErrorPaid                             = "%s is already paid"
	StatusErrorMiminumCharacter                 = "%s's minimum total characters is %s"
)

const (
	FormatEmailAdCreatedFree = "Dear %s, iklan anda sudah dibuat dengan nomor iklan [%s].\n" +
		"Berikut adalah password untuk mengaktifkan iklan anda agar dapat ditayangkan [%s].\n" +
		"Silahkan verifikasi iklan anda melalui fitur verifikasi di MobilCantik.\n\n" +

		"Gunakan password dan nomor iklan untuk melakukan penayangan, penutupan, maupun penghapusan iklan anda.\n\n" +
		"Salam, MobilCantik"

	FormatEmailAdCreatedPaid = "Dear %s, iklan anda sudah dibuat dengan nomor iklan [%s].\n" +
		"Berikut adalah password untuk mengaktifkan iklan anda agar dapat ditayangkan [%s].\n" +
		"Berikut adalah kode pembayaran untuk mengaktifkan iklan anda agar dapat ditayangkan [%s]. Silahkan lakukan pembayaran sebesar %s ke rekening yang telah disediakan di MobilCantik.\n\n" +

		"Gunakan password dan nomor iklan untuk melakukan konfirmasi pembayaran, penutupan, maupun penghapusan iklan anda.\n\n" +
		"Salam, MobilCantik"

	FormatEmailSuccessfullyConfirmed = "Dear %s, iklan anda dengan nomor iklan [%s] sudah berhasil ditayangkan.\n" +
		"Terima kasih sudah menggunakan jasa iklan MobilCantik.\n\n" +
		"Salam, MobilCantik"
)
