/**
 * PT Logam Gold Mulia Tbk — Main JavaScript
 * Handles: mobile menu, sticky header, scroll reveal, FAQ toggles, form validation
 */

(function () {
    'use strict';

    // =============================================
    // Mobile Menu
    // =============================================
    const menuBtn = document.getElementById('mobile-menu-btn');
    const mobileMenu = document.getElementById('mobile-menu');
    const iconOpen = document.getElementById('menu-icon-open');
    const iconClose = document.getElementById('menu-icon-close');

    if (menuBtn && mobileMenu) {
        menuBtn.addEventListener('click', function () {
            const isOpen = !mobileMenu.classList.contains('hidden');
            mobileMenu.classList.toggle('hidden');
            iconOpen.classList.toggle('hidden');
            iconClose.classList.toggle('hidden');
            menuBtn.setAttribute('aria-expanded', !isOpen);
        });
    }

    // =============================================
    // Sticky Header
    // =============================================
    const header = document.getElementById('site-header');
    if (header) {
        let lastScroll = 0;
        window.addEventListener('scroll', function () {
            const currentScroll = window.pageYOffset;
            if (currentScroll > 10) {
                header.classList.add('scrolled');
            } else {
                header.classList.remove('scrolled');
            }
            lastScroll = currentScroll;
        }, { passive: true });
    }

    // =============================================
    // Scroll Reveal
    // =============================================
    const revealElements = document.querySelectorAll('.scroll-reveal');
    if (revealElements.length > 0) {
        const revealObserver = new IntersectionObserver(function (entries) {
            entries.forEach(function (entry) {
                if (entry.isIntersecting) {
                    entry.target.classList.add('revealed');
                    revealObserver.unobserve(entry.target);
                }
            });
        }, {
            threshold: 0.1,
            rootMargin: '0px 0px -40px 0px'
        });

        revealElements.forEach(function (el) {
            revealObserver.observe(el);
        });
    }

    // =============================================
    // FAQ Toggles
    // =============================================
    const faqToggles = document.querySelectorAll('.faq-toggle');
    faqToggles.forEach(function (toggle) {
        toggle.addEventListener('click', function () {
            const item = this.closest('.faq-item');
            const content = item.querySelector('.faq-content');
            const isActive = item.classList.contains('active');

            // Close all others
            document.querySelectorAll('.faq-item.active').forEach(function (activeItem) {
                if (activeItem !== item) {
                    activeItem.classList.remove('active');
                    const activeContent = activeItem.querySelector('.faq-content');
                    activeContent.classList.add('hidden');
                    activeContent.style.maxHeight = null;
                    activeItem.querySelector('.faq-toggle').setAttribute('aria-expanded', 'false');
                }
            });

            // Toggle current
            if (isActive) {
                item.classList.remove('active');
                content.classList.add('hidden');
                content.style.maxHeight = null;
                this.setAttribute('aria-expanded', 'false');
            } else {
                item.classList.add('active');
                content.classList.remove('hidden');
                content.style.maxHeight = content.scrollHeight + 'px';
                this.setAttribute('aria-expanded', 'true');
            }
        });
    });

    // =============================================
    // Contact Form — Client-side Web3Forms Submit
    // =============================================
    const contactForm = document.getElementById('contact-form');
    if (contactForm) {
        contactForm.addEventListener('submit', function (e) {
            e.preventDefault();

            let isValid = true;
            const requiredFields = contactForm.querySelectorAll('[required]');

            // Clear previous error states
            contactForm.querySelectorAll('.field-error').forEach(function (el) {
                el.remove();
            });
            contactForm.querySelectorAll('.border-red-400').forEach(function (el) {
                el.classList.remove('border-red-400');
            });

            requiredFields.forEach(function (field) {
                var value = field.value.trim();
                if (!value) {
                    isValid = false;
                    showFieldError(field, 'Kolom ini wajib diisi');
                } else if (field.type === 'email' && !isValidEmail(value)) {
                    isValid = false;
                    showFieldError(field, 'Format email tidak valid');
                } else if (field.minLength && value.length < field.minLength) {
                    isValid = false;
                    showFieldError(field, 'Minimal ' + field.minLength + ' karakter');
                }
            });

            if (!isValid) return;

            // Hide previous messages
            var successEl = document.getElementById('form-success');
            var errorEl = document.getElementById('form-error');
            if (successEl) successEl.classList.add('hidden');
            if (successEl) successEl.classList.remove('flex');
            if (errorEl) errorEl.classList.add('hidden');
            if (errorEl) errorEl.classList.remove('flex');

            // Show loading state
            var submitBtn = document.getElementById('submit-btn');
            var btnText = document.getElementById('btn-text');
            var btnIcon = document.getElementById('btn-icon');
            var btnSpinner = document.getElementById('btn-spinner');
            if (submitBtn) submitBtn.disabled = true;
            if (btnText) btnText.textContent = 'Mengirim...';
            if (btnIcon) btnIcon.classList.add('hidden');
            if (btnSpinner) btnSpinner.classList.remove('hidden');

            var formData = new FormData(contactForm);
            var object = {};
            formData.forEach(function (value, key) {
                object[key] = value;
            });

            fetch('https://api.web3forms.com/submit', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json'
                },
                body: JSON.stringify(object)
            })
            .then(function (response) {
                return response.json();
            })
            .then(function (data) {
                if (data.success) {
                    if (successEl) {
                        successEl.classList.remove('hidden');
                        successEl.classList.add('flex');
                    }
                    contactForm.reset();
                } else {
                    var errorTextEl = document.getElementById('form-error-text');
                    if (errorTextEl) errorTextEl.textContent = data.message || 'Maaf, terjadi kesalahan saat mengirim pesan.';
                    if (errorEl) {
                        errorEl.classList.remove('hidden');
                        errorEl.classList.add('flex');
                    }
                }
            })
            .catch(function () {
                if (errorEl) {
                    errorEl.classList.remove('hidden');
                    errorEl.classList.add('flex');
                }
            })
            .finally(function () {
                if (submitBtn) submitBtn.disabled = false;
                if (btnText) btnText.textContent = 'Kirim Pesan';
                if (btnIcon) btnIcon.classList.remove('hidden');
                if (btnSpinner) btnSpinner.classList.add('hidden');
            });
        });
    }

    function showFieldError(field, message) {
        field.classList.add('border-red-400');
        const errorEl = document.createElement('p');
        errorEl.className = 'field-error text-xs text-red-600 mt-1';
        errorEl.textContent = message;
        field.parentNode.appendChild(errorEl);
    }

    function isValidEmail(email) {
        return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
    }
})();
